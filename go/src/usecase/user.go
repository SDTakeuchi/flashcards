package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/adapter/config"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"github.com/SDTakeuchi/go/src/flashcards/domain/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserUsecase interface {
	SignUp(ctx context.Context, name string, password *model.Password) error
	Login(ctx context.Context, in *LoginInput) (*LoginOutput, error)
}

type userUsecase struct {
	userRepo    repo.UserRepo
	sessionRepo repo.SessionRepo
	conn        *gorm.DB
	tokenIssuer auth.TokenIssuer
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(
	userRepo repo.UserRepo,
	sessionRepo repo.SessionRepo,
	conn *gorm.DB,
	tokenIssuer auth.TokenIssuer,
) *userUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		conn:        conn,
		tokenIssuer: tokenIssuer,
	}
}

func (u *userUsecase) SignUp(ctx context.Context, name string, password *model.Password) error {
	if password == nil {
		return errors.New("password is empty")
	}

	_, err := u.userRepo.GetUser(ctx, u.conn, name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			return errors.New("user already exists")
		}
		return err
	}

	user, err := model.NewUser(name, *password)
	if err != nil {
		return err
	}

	return u.userRepo.Create(ctx, u.conn, user)
}

type (
	LoginInput struct {
		Name        string
		RawPassword string
		ClientIP    string
		UserAgent   string
	}
	LoginOutput struct {
		AccessToken           string
		AccessTokenExpiresAt  time.Time
		RefreshToken          string
		RefreshTokenExpiresAt time.Time
		UserID                uuid.UUID
	}
)

func (u *userUsecase) Login(ctx context.Context, in *LoginInput) (*LoginOutput, error) {
	user, err := u.userRepo.GetUser(ctx, u.conn, in.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err = model.IsEqual(user.Password(), in.RawPassword); err != nil {
		return nil, err
	}

	// issue token
	accessToken, accessPayload, err := u.tokenIssuer.Create(
		user.ID(),
		auth.AccessToken,
		config.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := u.tokenIssuer.Create(
		user.ID(),
		auth.RefreshToken,
		config.Get().Token.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	if err = u.createSession(
		ctx,
		refreshToken,
		*refreshPayload,
		in.ClientIP,
		in.UserAgent,
	); err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt(),
		UserID:                user.ID(),
	}, nil
}

func (u *userUsecase) createSession(
	ctx context.Context,
	refreshToken string,
	payload auth.Payload,
	clientIP string,
	userAgent string,
) error {
	// token is refresh token
	session, err := auth.NewSession(
		payload.ID(),
		payload.UserID(),
		refreshToken,
		userAgent,
		clientIP,
		payload.ExpiresAt(),
		payload.IssuedAt(),
	)
	if err != nil {
		return err
	}

	return u.sessionRepo.Create(ctx, u.conn, session)
}

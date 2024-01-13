package usecase

import (
	"context"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
)

type UserUsecase interface {
	SignUp(ctx context.Context, name, password *model.Password) error
	Login(ctx context.Context, name, password *model.Password) error
}

type userUsecase struct{}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase() *userUsecase {
	return &userUsecase{}
}

func (u *userUsecase) SignUp(ctx context.Context, name, password *model.Password) error { return nil }
func (u *userUsecase) Login(ctx context.Context, name, password *model.Password) error  { return nil }

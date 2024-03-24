package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/adapter/config"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtMethod = jwt.SigningMethodHS256

type JWTIssuer struct {
	secretKey string
}

func NewJWTIssuer(secretKey string) (auth.TokenIssuer, error) {
	minSecretKeySize := config.Get().Token.MinSecretKeySize

	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTIssuer{secretKey: secretKey}, nil
}

func (issuer *JWTIssuer) Create(
	userID uuid.UUID,
	tokenType auth.TokenType,
	duration time.Duration,
) (string, *auth.Payload, error) {
	// test code cannot read config, so provide default value here
	if duration == time.Duration(0) {
		duration = time.Minute
	}
	payload := auth.NewPayload(userID, tokenType, duration)
	claims := ClaimsFromPayload(*payload)
	token := jwt.NewWithClaims(jwtMethod, claims)
	signedToken, err := token.SignedString([]byte(issuer.secretKey))
	return signedToken, payload, err
}

func (issuer *JWTIssuer) Verify(token string) (*auth.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrTokenInvalid
		}
		return []byte(issuer.secretKey), nil
	}

	parsedToken, err := jwt.Parse(
		token,
		keyFunc,
		jwt.WithValidMethods([]string{"HS256"}),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, auth.ErrTokenExpired
		}
		return nil, auth.ErrTokenInvalid
	}
	if !parsedToken.Valid {
		return nil, auth.ErrTokenInvalid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, auth.ErrTokenInvalid
	}

	payload, err := PayloadFromClaims(claims)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

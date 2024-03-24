package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	// errors returned by VerifyToken method
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("token is invalid")
)

type TokenType int

const (
	AccessToken TokenType = iota + 1
	RefreshToken
)

type Payload struct {
	id        uuid.UUID
	tokenType TokenType
	userID    uuid.UUID
	issuedAt  time.Time
	expiresAt time.Time
}

func (p *Payload) ID() uuid.UUID        { return p.id }
func (p *Payload) TokenType() TokenType { return p.tokenType }
func (p *Payload) UserID() uuid.UUID    { return p.userID }
func (p *Payload) IssuedAt() time.Time  { return p.issuedAt }
func (p *Payload) ExpiresAt() time.Time { return p.expiresAt }

func NewPayload(userID uuid.UUID, tokenType TokenType, duration time.Duration) *Payload {
	return &Payload{
		id:        uuid.New(),
		tokenType: tokenType,
		userID:    userID,
		issuedAt:  time.Now(),
		expiresAt: time.Now().Add(duration),
	}
}

func PayloadFromPersistance(
	id uuid.UUID,
	tokenType TokenType,
	userID uuid.UUID,
	issuedAt time.Time,
	expiresAt time.Time,
) *Payload {
	return &Payload{
		id: id,
		tokenType: tokenType,
		userID: userID,
		issuedAt: issuedAt,
		expiresAt: expiresAt,
	}
}

func ClaimsFromPayload(p Payload) jwt.MapClaims {
	return jwt.MapClaims{
		"id":   p.ID(),
		"type": p.TokenType(),
		"sub":  p.UserID(),
		"exp":  p.ExpiresAt().Unix(),
		"iat":  p.IssuedAt().Unix(),
	}
}

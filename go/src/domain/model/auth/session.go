package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	id           uuid.UUID // ID originally comes from refresh token's payload id
	userID       uuid.UUID
	refreshToken string
	userAgent    string
	clientIP     string
	expiresAt    time.Time
	createdAt    time.Time
}

func (s *Session) ID() uuid.UUID        { return s.id }
func (s *Session) UserID() uuid.UUID    { return s.userID }
func (s *Session) RefreshToken() string { return s.refreshToken }
func (s *Session) UserAgent() string    { return s.userAgent }
func (s *Session) ClientIP() string     { return s.clientIP }
func (s *Session) ExpiresAt() time.Time { return s.expiresAt }
func (s *Session) CreatedAt() time.Time { return s.createdAt }

func validateSession(s Session) error {
	if s.RefreshToken() == "" {
		return errors.New("session must have a refresh token string")
	}
	if s.ExpiresAt().IsZero() {
		return errors.New("session must have expiration date")
	}
	return nil
}

func NewSession(
	id uuid.UUID,
	userID uuid.UUID,
	refreshToken string,
	userAgent string,
	clientIP string,
	expiresAt time.Time,
	createdAt time.Time,
) (*Session, error) {
	s := &Session{
		id:           id,
		userID:       userID,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		clientIP:     clientIP,
		expiresAt:    expiresAt,
		createdAt:    createdAt,
	}
	if err := validateSession(*s); err != nil {
		return nil, err
	}
	return s, nil
}

func SessionFromPersistance(
	id uuid.UUID,
	userID uuid.UUID,
	refreshToken string,
	userAgent string,
	clientIP string,
	expiresAt time.Time,
	createdAt time.Time,
) *Session {
	return &Session{
		id:           id,
		userID:       userID,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		clientIP:     clientIP,
		expiresAt:    expiresAt,
		createdAt:    createdAt,
	}
}

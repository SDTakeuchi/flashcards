package model

import (
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	RefreshToken string
	UserAgent    string
	ClientIP     string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

func SessionFromDomainModel(s *auth.Session) *Session {
	return &Session{
		ID:           s.ID(),
		UserID:       s.UserID(),
		RefreshToken: s.RefreshToken(),
		UserAgent:    s.ClientIP(),
		ClientIP:     s.UserAgent(),
		ExpiresAt:    s.ExpiresAt(),
		CreatedAt:    s.CreatedAt(),
	}
}

package repo

import (
	"context"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"gorm.io/gorm"
)

type SessionRepo interface {
	Create(ctx context.Context, conn *gorm.DB, s *auth.Session) error
	GetByID(ctx context.Context, conn *gorm.DB, id string) (*auth.Session, error)
}

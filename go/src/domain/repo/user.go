package repo

import (
	"context"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, conn *gorm.DB, u *model.User) error
	GetUser(ctx context.Context, conn *gorm.DB, name string) (*model.User, error)
}

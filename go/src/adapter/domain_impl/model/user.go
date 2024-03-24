package model

import (
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Name           string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func UserFromDomainModel(u *model.User) *User {
	return &User{
		ID:             u.ID(),
		Name:           u.Name(),
		HashedPassword: u.Password().String(),
		CreatedAt:      u.CreatedAt(),
		UpdatedAt:      u.UpdatedAt(),
	}
}

package repo

import (
	"context"

	modelimpl "github.com/SDTakeuchi/go/src/flashcards/adapter/domain_impl/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/repo"
	"gorm.io/gorm"
)

type UserRepo struct{}

var _ repo.UserRepo = (*UserRepo)(nil)

func NewUserRepo() repo.UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Create(ctx context.Context, conn *gorm.DB, u *model.User) error {
	userImpl := modelimpl.UserFromDomainModel(u)
	return conn.Create(userImpl).Error
}

func (r *UserRepo) GetUser(ctx context.Context, conn *gorm.DB, name string) (*model.User, error) {
	u := new(modelimpl.User)

	if err := conn.Where(&modelimpl.User{Name: name}).Take(u).Error; err != nil {
		return nil, err
	}

	return model.UserFromPersistence(
		u.ID.String(),
		u.Name,
		u.HashedPassword,
		u.CreatedAt,
		u.UpdatedAt,
	), nil
}

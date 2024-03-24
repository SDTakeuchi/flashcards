package repo

import (
	"context"

	modelimpl "github.com/SDTakeuchi/go/src/flashcards/adapter/domain_impl/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"github.com/SDTakeuchi/go/src/flashcards/domain/repo"
	"gorm.io/gorm"
)

type SessionRepo struct{}

var _ repo.SessionRepo = (*SessionRepo)(nil)

func NewSessionRepo() repo.SessionRepo {
	return &SessionRepo{}
}

func (r *SessionRepo) Create(ctx context.Context, conn *gorm.DB, s *auth.Session) error {
	session := modelimpl.SessionFromDomainModel(s)
	return conn.Create(session).Error
}

func (r *SessionRepo) GetByID(ctx context.Context, conn *gorm.DB, id string) (*auth.Session, error) {
	s := new(auth.Session)
	if err := conn.Where("id = ?", id).Take(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

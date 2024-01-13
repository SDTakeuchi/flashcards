package usecase

import (
	"context"

	"github.com/SDTakeuchi/flashcards/domain/model"
	"github.com/SDTakeuchi/flashcards/domain/repo"
	"github.com/google/uuid"
)

type CardUsecase interface {
	GetOldest(ctx context.Context) (*model.Card, error)
	Create(ctx context.Context, word, description string, userID uuid.UUID) error
	Update(ctx context.Context, word, description string, status model.CardStatus) error
}

type cardUsecase struct {
	cardRepo repo.CardRepo
}

var _ CardUsecase = (*cardUsecase)(nil)

func NewCardUsecase(cardRepo repo.CardRepo) *cardUsecase {
	return &cardUsecase{cardRepo: cardRepo}
}

func (u *cardUsecase) GetOldest(ctx context.Context) (*model.Card, error) {
	return u.cardRepo.GetLastUpdated()
}

func (u *cardUsecase) Create(ctx context.Context, word, description string, userID uuid.UUID) error {
	return nil
}

func (u *cardUsecase) Update(ctx context.Context, word, description string, status model.CardStatus) error {
	c, row, err := u.cardRepo.GetByWordDescription(word, description)
	if err != nil {
		return err
	}
	c.Seen()
	c.SetStatus(status)
	return u.cardRepo.UpdateStatus(row, c)
}

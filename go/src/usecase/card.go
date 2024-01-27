package usecase

import (
	"context"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/repo"
	"github.com/google/uuid"
)

type CardUsecase interface {
	GetOldest(ctx context.Context) (*model.Card, error)
	GetRemembered(ctx context.Context) (*model.Card, error)
	Create(ctx context.Context, word, description string, userID uuid.UUID) error
	Update(ctx context.Context, word, description string, status model.CardStatus) error
}

type CardOutput struct {
	Meta *Meta       `json:"meta"`
	Card *model.Card `json:"card"`
}

type Meta struct {
	Total              int `json:"total"`
	TotalSeenToday     int `json:"total_seen_today"`
	TotalUnspecified   int `json:"total_unspecified"`
	TotalRemembered    int `json:"total_remembered"`
	TotalLearnAgain    int `json:"total_learn_again"`
	TotalNotRemembered int `json:"total_not_remembered"`
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

func (u *cardUsecase) GetRemembered(ctx context.Context) (*model.Card, error) {
	return u.cardRepo.GetLastRemembered()
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

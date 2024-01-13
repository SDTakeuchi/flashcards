package repo

import (
	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
)

type CardRepo interface {
	GetByWordDescription(word, description string) (*model.Card, uint, error)
	GetLastUpdated() (*model.Card, error)
	UpdateStatus(row uint, c *model.Card) error
}

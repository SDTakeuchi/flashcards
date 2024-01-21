package handler

import (
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
)

type TimeMixin struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Card struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ID            string    `json:"id"`
	Word          string    `json:"word"`
	Description   string    `json:"description"`
	LastSeen      time.Time `json:"last_seen,omitempty"`
	PartOfSpeech  string    `json:"part_of_speech"`
	Example       string    `json:"example"`
	Pronunciation string    `json:"pronunciation"`
	Status        string    `json:"status"`
	UserID        string    `json:"user_id"`
}

func convertCard(m *model.Card) *Card {
	return &Card{
		CreatedAt:     m.CreatedAt(),
		UpdatedAt:     m.UpdatedAt(),
		ID:            m.ID().String(),
		Word:          m.Word(),
		Description:   m.Description(),
		LastSeen:      *m.LastSeen(),
		PartOfSpeech:  m.PartOfSpeech().String(),
		Example:       m.Example(),
		Pronunciation: m.Pronunciation(),
		Status:        m.Status().String(),
		UserID:        m.UserID().String(),
	}
}

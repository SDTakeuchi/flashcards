package repo

import (
	"errors"
	"strconv"
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/adapter/config"
	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
	"github.com/SDTakeuchi/go/src/flashcards/domain/repo"
	"github.com/SDTakeuchi/go/src/flashcards/pkg/google/spreadsheet"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("item not found")

type cardRepo struct {
	sheetService spreadsheet.SpreadsheetService
}

var _ repo.CardRepo = (*cardRepo)(nil)

func NewCardRepo(sheetService spreadsheet.SpreadsheetService) repo.CardRepo {
	return &cardRepo{sheetService: sheetService}
}

func (r *cardRepo) GetByWordDescription(word, description string) (*model.Card, uint, error) {
	values, err := r.sheetService.Read(config.Get().SheetID)
	if err != nil {
		return nil, 0, err
	}

	for i, v := range values {
		if v == nil {
			continue
		}
		if v.Word == word && v.Description == description {
			return model.CardFromPersistence(
				uuid.NewString(),
				v.Word,
				v.Description,
				uuid.NewString(),
				uint8(v.Status),
				&v.LastSeen,
				v.PartOfSpeech,
				v.Example,
				v.CreatedAt,
				v.UpdatedAt,
			), uint(i + 2), nil
		}
	}
	return nil, 0, ErrNotFound
}

func (r *cardRepo) GetLastUpdated() (*model.Card, error) {
	values, err := r.sheetService.Read(config.Get().SheetID)
	if err != nil {
		return nil, err
	}

	var MostLastSeen *spreadsheet.Value
	for _, v := range values {
		if v == nil {
			continue
		}
		if v.Status == int(model.CardStatusRemembered) {
			continue
		}
		if MostLastSeen == nil || v.LastSeen.Before(MostLastSeen.LastSeen) {
			MostLastSeen = v
		}
	}

	return model.CardFromPersistence(
		uuid.NewString(),
		MostLastSeen.Word,
		MostLastSeen.Description,
		uuid.NewString(),
		uint8(MostLastSeen.Status),
		&MostLastSeen.LastSeen,
		MostLastSeen.PartOfSpeech,
		MostLastSeen.Example,
		MostLastSeen.CreatedAt,
		MostLastSeen.UpdatedAt,
	), nil
}

func (r *cardRepo) UpdateStatus(row uint, c *model.Card) error {
	return r.sheetService.Update(
		config.Get().SheetID,
		row,
		c.Word(),
		c.Description(),
		c.PartOfSpeech().String(),
		c.Example(),
		c.LastSeen().Format(time.RFC3339),
		strconv.Itoa(int(c.Status().Uint8())),
		c.CreatedAt().Format(time.RFC3339),
		c.UpdatedAt().Format(time.RFC3339),
	)
}

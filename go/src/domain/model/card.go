package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	TimeMixin
	id            uuid.UUID
	word          string
	description   string
	lastSeen      *time.Time
	status        CardStatus
	userID        uuid.UUID
	partOfSpeech  PartOfSpeech
	example       string
	pronunciation string
}

func NewCard(word, description, example, pronunciation string, pos PartOfSpeech, userID *uuid.UUID) (*Card, error) {
	c := &Card{
		id:            uuid.New(),
		word:          word,
		description:   description,
		status:        CardStatusUnspecified,
		lastSeen:      nil,
		partOfSpeech:  pos,
		example:       example,
		pronunciation: pronunciation,
		TimeMixin:     *NewTimeMixin(),
	}
	if userID != nil {
		c.userID = *userID
	}
	if err := validateCard(c); err != nil {
		return nil, err
	}
	return c, nil
}

var ErrFailedToValidateCard error = errors.New("failed to validate card")

func validateCard(c *Card) error {
	if c.word == "" {
		return fmt.Errorf("%w: word is empty", ErrFailedToValidateCard)
	}
	if c.description == "" {
		return fmt.Errorf("%w: description is empty", ErrFailedToValidateCard)
	}
	if c.userID == uuid.Nil {
		return fmt.Errorf("%w: userID is empty", ErrFailedToValidateCard)
	}
	return nil
}

func CardFromPersistence(
	id,
	word,
	description,
	userID string,
	status uint8,
	lastSeen *time.Time,
	pos,
	example,
	pronunciation string,
	createdAt,
	updatedAt time.Time,
) *Card {
	return &Card{
		id:            uuid.MustParse(id),
		word:          word,
		description:   description,
		status:        CardStatusFromUint8(status),
		lastSeen:      lastSeen,
		userID:        uuid.MustParse(userID),
		partOfSpeech:  PartOfSpeechFromString(pos),
		example:       example,
		pronunciation: pronunciation,
		TimeMixin:     *TimeMixinFromPersistence(createdAt, updatedAt),
	}
}

func (c Card) ID() uuid.UUID              { return c.id }
func (c Card) Word() string               { return c.word }
func (c Card) Description() string        { return c.description }
func (c Card) LastSeen() *time.Time       { return c.lastSeen }
func (c Card) Status() CardStatus         { return c.status }
func (c Card) PartOfSpeech() PartOfSpeech { return c.partOfSpeech }
func (c Card) Example() string            { return c.example }
func (c Card) Pronunciation() string      { return c.pronunciation }
func (c Card) UserID() uuid.UUID          { return c.userID }

func (c *Card) Seen() {
	now := time.Now()
	c.lastSeen = &now
	c.Updated()
}

func (c *Card) SetStatus(status CardStatus) {
	c.status = status
	c.Updated()
}

func (c *Card) SetWord(word string) error {
	c.word = word
	if err := validateCard(c); err != nil {
		return err
	}
	c.Updated()
	return nil
}

func (c *Card) SetDescription(description string) error {
	c.description = description
	if err := validateCard(c); err != nil {
		return err
	}
	c.Updated()
	return nil
}

package model

import "time"

type TimeMixin struct {
	createdAt time.Time
	updatedAt time.Time
}

func NewTimeMixin() *TimeMixin {
	return &TimeMixin{
		createdAt: time.Now(),
		updatedAt: time.Time{},
	}
}

func TimeMixinFromPersistence(createdAt, updatedAt time.Time) *TimeMixin {
	return &TimeMixin{
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (t TimeMixin) CreatedAt() time.Time { return t.createdAt }
func (t TimeMixin) UpdatedAt() time.Time { return t.updatedAt }

func (t *TimeMixin) Updated() {
	t.updatedAt = time.Now()
}

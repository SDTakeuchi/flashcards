package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	TimeMixin
	id       uuid.UUID
	name     string
	password Password
}

func NewUser(name string, password Password) (*User, error) {
	u := &User{
		id:        uuid.New(),
		name:      name,
		password:  password,
		TimeMixin: *NewTimeMixin(),
	}
	if err := validateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

var ErrFailedToValidateUser error = errors.New("failed to validate user")

func validateUser(c *User) error {
	if c.name == "" {
		return fmt.Errorf("%w: name is empty", ErrFailedToValidateUser)
	}
	if &c.password == nil {
		return fmt.Errorf("%w: password is empty", ErrFailedToValidateUser)
	}
	return nil
}

func UserFromPersistence(id, name, password string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        uuid.MustParse(id),
		name:      name,
		password:  *PasswordFromPersistence(password),
		TimeMixin: *TimeMixinFromPersistence(createdAt, updatedAt),
	}
}

func (u *User) ID() uuid.UUID      { return u.id }
func (u *User) Password() Password { return u.password }
func (u *User) Name() string       { return u.name }

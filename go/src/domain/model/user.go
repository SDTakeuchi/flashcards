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

func NewUser(name string, description string, password Password) (*User, error) {
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

func (c *User) ID() uuid.UUID { return c.id }
func (c *User) Name() string  { return c.name }

// TODO: implement update
// func (c *User) SetName(name string) {
// 	c.name = name
//     c.Updated()
// }

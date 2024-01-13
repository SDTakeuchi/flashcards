package model

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	minPasswordLength uint8 = 8

	passwordsBlacklist = []string{
		"password",
		"12345678",
		"123456789",
		"qwertyui",
		"asdfghjk",
	}

	ErrPasswordTooSimple = errors.New("password is too simple")
	ErrPasswordTooShort  = fmt.Errorf("password must be at least %d characters", minPasswordLength)
	ErrIncorrectPassword = errors.New("wrong password")
)

func MinPasswordLength() uint8     { return minPasswordLength }
func SetMinPasswordLength(l uint8) { minPasswordLength = l }

func AppendPasswordToBlacklist(p ...string) {
	passwordsBlacklist = append(passwordsBlacklist, p...)
}

// Password is a value object that represents a hashed password
type Password string

func (p Password) String() string { return string(p) }

func NewPassword(raw string) (*Password, error) {
	err := validatePassword(raw)
	if err != nil {
		return nil, err
	}
	var hashed string
	hashed, err = hash(raw)
	if err != nil {
		return nil, err
	}
	p := Password(hashed)
	return &p, nil
}

func PasswordFromPersistence(s string) *Password {
	p := Password(s)
	return &p
}

// return no error if hashed password and incoming password match
func IsEqual(p Password, incoming string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(p.String()),
		[]byte(incoming),
	); err != nil {
		return ErrIncorrectPassword
	}
	return nil
}

func validatePassword(p string) error {
	if len(p) < int(minPasswordLength) {
		return ErrPasswordTooShort
	}
	for _, sp := range passwordsBlacklist {
		if p == sp {
			return ErrPasswordTooSimple
		}
	}
	return nil
}

func hash(p string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashed), nil
}

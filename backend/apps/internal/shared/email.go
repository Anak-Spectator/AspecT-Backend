package shared

import (
	"errors"
	"regexp"
)

var emailFormatValidator = regexp.MustCompile("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$")

type Email string

func NewEmail(email string) (Email, error) {
	e := Email(email)
	if err := e.Validate(); err != nil {
		return "", err
	}

	return e, nil
}

func (email Email) Validate() error {
	if email == "" {
		return errors.New("Email can't be empty")
	}

	if !emailFormatValidator.Match([]byte(email)) {
		return errors.New("Invalid email format")
	}
	return nil
}

package account

import (
	"aspect_apps/internal/shared"
	"errors"
	"log"
)

type Password string

func NewPassword(password string) (Password, error) {
	pwd := Password(password)

	if err := pwd.Validate(); err != nil {
		return "", err
	}

	return pwd, nil
}

func (password Password) Validate() error {
	if password == "" {
		return errors.New("password can't be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (password Password) Hash(hasher Hasher) (HashedPassword, error) {
	hashedPassword, err := hasher.Hash(string(password))
	if err != nil {
		if errors.Is(err, &shared.InfrastructureError{}) {
			log.Println(err)
		}
		return "", err
	}

	return HashedPassword(hashedPassword), nil
}

type HashedPassword string

func (hashedPassword HashedPassword) Match(hasher Hasher, password Password) bool {
	return hasher.IsInputHashSimilar(string(password), string(hashedPassword))
}

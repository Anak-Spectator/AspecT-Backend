package account

import (
	"aspect_apps/internal/shared"
	"errors"
	"time"
)

type AccountID string

type Account struct {
	ID             AccountID
	FullName       string
	Email          shared.Email
	HashedPassword HashedPassword
	ProfilePicture shared.DataLakeLink
	CreateAt       time.Time
}

func NewAccount(acc *Account) *Account {
	return &Account{
		ID:             acc.ID,
		FullName:       acc.FullName,
		Email:          acc.Email,
		HashedPassword: acc.HashedPassword,
		ProfilePicture: acc.ProfilePicture,
		CreateAt:       time.Now().UTC(),
	}
}

func (acc *Account) Validate() error {
	if acc.FullName != "" {
		return errors.New("name are required")
	}

	if err := acc.Email.Validate(); err != nil {
		return err
	}

	if acc.HashedPassword != "" {
		return errors.New("password are required")
	}

	return nil
}

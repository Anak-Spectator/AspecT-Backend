package account

import (
	"aspect_apps/internal/shared"
	"time"
)

type AccountFactory struct {
	idGenerator shared.IDGenerator
	hasher      shared.Hasher
}

func NewAccountFactory(idGenerator shared.IDGenerator, hasher shared.Hasher) *AccountFactory {
	return &AccountFactory{
		idGenerator: idGenerator,
		hasher:      hasher,
	}
}

func (factory *AccountFactory) AddUser(email, password, fullname string) (*Account, error) {
	id, err := factory.idGenerator.GenerateID()
	if err != nil {
		return &Account{}, err
	}
	rawPassword, err := NewPassword(password)

	if err != nil {
		return &Account{}, err
	}
	hasedPassword, err := rawPassword.Hash(factory.hasher)

	user := &Account{
		ID:             AccountID(id),
		FullName:       fullname,
		Email:          shared.Email(email),
		HashedPassword: hasedPassword,
		CreateAt:       time.Now().UTC(),
	}

	return user, nil
}

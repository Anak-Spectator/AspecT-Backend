package accountquery

import (
	"aspect_apps/internal/services/account"
	"aspect_apps/internal/shared"
)

type gormAccountInfo struct {
	ID             account.AccountID `gorm:"primary_key;unique;not null;unique_index"`
	FullName       string            `gorm:"index"`
	Email          shared.Email      `gorm:"index;unique;not null;unique_index"`
	ProfilePicture shared.DataLakeLink
	HashedPassword account.HashedPassword
	Version        shared.AggregateVersion
}

func newGormAccount(acc *account.Account) *gormAccountInfo {
	return &gormAccountInfo{
		ID:             acc.ID,
		FullName:       acc.FullName,
		Email:          acc.Email,
		ProfilePicture: acc.ProfilePicture,
		HashedPassword: acc.HashedPassword,
	}
}

func (acc *gormAccountInfo) ToAccountModel() *account.Account {
	return &account.Account{
		ID:             acc.ID,
		FullName:       acc.FullName,
		Email:          acc.Email,
		ProfilePicture: acc.ProfilePicture,
		HashedPassword: acc.HashedPassword,
	}
}

func (acc *gormAccountInfo) TableName() string {
	return "account_user"
}

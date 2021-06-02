package account

import "aspect_apps/internal/shared"

type AccountRepository interface {
	FindByID(id AccountID) (*Account, shared.AggregateVersion, error)
	FindByEmail(email shared.Email) (*Account, shared.AggregateVersion, error)
	Save(*Account, shared.AggregateVersion) (*Account, error)
}

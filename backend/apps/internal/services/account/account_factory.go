package account

import (
	"aspect_apps/internal/services/account/accountinfra"
	"aspect_apps/internal/shared"
)

type AccountFactory struct {
	idGenerator shared.IDGenerator
	bycrypt     accountinfra.BcryptHasher
}

package accountapp

import (
	"aspect_apps/internal/services/account"
)

type AccountApplicationService struct {
	accRepo account.AccountRepository
	// accQuery *account.AccountQuery
}

func NewAccountApplicationService(accRepo account.AccountRepository) *AccountApplicationService {
	return &AccountApplicationService{
		accRepo: accRepo,
	}
}

func (svc *AccountApplicationService) Test(str string) string {
	return str
}

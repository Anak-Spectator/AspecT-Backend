package accountapp

import (
	"aspect_apps/internal/services/account"
)

type AccountApplicationService struct {
	accRepo    account.AccountRepository
	accFactory account.AccountFactory
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

func (svc *AccountApplicationService) RegisterNewUser() {

}

func (svc *AccountApplicationService) EditUserProfile() {
}

func (svc *AccountApplicationService) UploadProfileImage() {
}

func (svc *AccountApplicationService) DeleteProfileImage() {

}

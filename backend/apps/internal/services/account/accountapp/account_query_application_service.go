package accountapp

import "aspect_apps/internal/services/account"

type AccountQueryApplicationService struct {
	accQuery *account.AccountQuery
}

func NewAccountQueryApplicationService(accQuery *account.AccountQuery) *AccountQueryApplicationService {
	return &AccountQueryApplicationService{
		accQuery: accQuery,
	}
}

func (svc *AccountQueryApplicationService) GetUserInfo() {

}

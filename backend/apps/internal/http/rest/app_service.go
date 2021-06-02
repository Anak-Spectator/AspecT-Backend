package rest

import "aspect_apps/internal/services/account/accountapp"

type AccountMainService struct {
	accAppSvc *accountapp.AccountApplicationService
}

func NewAccountMainService(accAppSvc *accountapp.AccountApplicationService) *AccountMainService {
	return &AccountMainService{
		accAppSvc: accAppSvc,
	}
}

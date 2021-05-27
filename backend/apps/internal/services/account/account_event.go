package account

import "time"

type AccountRegitered struct {
	Account   Account
	createdAt time.Time
}

func (event AccountRegitered) Name() string {
	return "account_user_registered"
}

func (event AccountRegitered) Data() interface{} {
	return event
}

func (event AccountRegitered) CreatedAt() time.Time {
	return event.createdAt
}

type UserUploadPictureProfile struct {
	AccountID AccountID
}

package account

type AccountQuery interface {
	Info(id AccountID) (*Account, error)
}

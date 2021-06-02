package shared

type UserIdentity interface{}

type UserIdentifier interface {
	Identify(identity UserIdentity) (UserID, error)
}

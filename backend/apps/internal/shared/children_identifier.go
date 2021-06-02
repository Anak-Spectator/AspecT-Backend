package shared

type ChildrenIdentity interface{}

type ChildrenIdentifier interface {
	Identify(identity ChildrenIdentity) (ChildrenID, error)
}

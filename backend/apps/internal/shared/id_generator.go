package shared

type ID string

type IDGenerator interface {
	GenerateID() (ID, error)
}

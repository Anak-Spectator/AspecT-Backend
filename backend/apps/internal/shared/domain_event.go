package shared

import "time"

type DomainEvent interface {
	Name() string
	Data() interface{}
	CreatedAt() time.Time
}

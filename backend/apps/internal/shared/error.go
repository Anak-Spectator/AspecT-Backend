package shared

import (
	"errors"
	"fmt"
)

//InfrastructureError spawned if there are error related to infrastructure
type InfrastructureError struct {
	err error
}

func NewInfrastructureError(err error) *InfrastructureError {
	return &InfrastructureError{
		err: err,
	}
}

func (err *InfrastructureError) Error() string {
	return fmt.Sprintf("There is infrastructure error which is: %v", err.err.Error())
}

var AggregateVersionOutdatedError = errors.New("Current version of aggregate is outdated")

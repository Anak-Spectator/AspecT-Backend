package uuid

import (
	"aspect_apps/internal/shared"
	"fmt"

	"github.com/google/uuid"
)

type idGenerator struct {
}

func NewIDGenerator() *idGenerator {
	return &idGenerator{}
}

func (generator *idGenerator) GenerateID() (shared.ID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("There is something wrong with id generator which is: %w", err)
	}

	return shared.ID(uuid.String()), nil
}

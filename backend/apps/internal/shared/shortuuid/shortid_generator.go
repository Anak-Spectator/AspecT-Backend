package shortuuid

import (
	"aspect_apps/internal/shared"
	"fmt"
	"regexp"
	"strings"

	"github.com/teris-io/shortid"
)

type idGenerator struct {
}

func NewIDGenerator() *idGenerator {
	return &idGenerator{}
}

func (generator *idGenerator) GenerateID() (shared.ID, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)

	if err != nil {
		return "", fmt.Errorf("There is something wrong with id generator which is: %w", err)
	}

	id, err := sid.Generate()
	if err != nil {
		return "", fmt.Errorf("There is something wrong when generate an id : %w", err)
	}

	re, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		return "", fmt.Errorf("There is something wrong with id generator : %w", err)
	}

	newID := re.ReplaceAllString(id, "")
	newID = strings.ToUpper(newID)
	return shared.ID(newID), nil
}

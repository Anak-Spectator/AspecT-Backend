package accountinfra

import (
	"aspect_apps/internal/shared"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (hasher *BcryptHasher) Hash(input string) (string, error) {
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", shared.NewInfrastructureError(err)
	}

	return string(hashedInput), nil
}

func (hasher *BcryptHasher) IsInputHashSimilar(input string, hashedInput string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedInput), []byte(input)) == nil
}

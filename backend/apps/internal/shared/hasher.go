package shared

type Hasher interface {
	Hash(input string) (string, error)
	IsInputHashSimilar(input string, hashedInput string) bool
}

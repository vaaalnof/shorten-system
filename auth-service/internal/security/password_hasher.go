package security

import "golang.org/x/crypto/bcrypt"

type PasswordHash interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, plainPassword string) error
}

type BcryptHash struct{}

func NewBcryptHash() *BcryptHash {
	return &BcryptHash{}
}

func (h *BcryptHash) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *BcryptHash) Compare(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

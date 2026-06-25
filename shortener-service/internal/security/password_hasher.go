package security

import "golang.org/x/crypto/bcrypt"

type PasswordHash interface {
	Hash(
		password string,
	) (
		string,
		error,
	)

	Compare(
		hashedPassword string,
		password string,
	) error
}

type BcryptHash struct {
	cost int
}

func NewPasswordHash() *BcryptHash {

	return &BcryptHash{
		cost: bcrypt.DefaultCost,
	}
}

func (h *BcryptHash) Hash(
	password string,
) (
	string,
	error,
) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		h.cost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *BcryptHash) Compare(
	hashedPassword string,
	password string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

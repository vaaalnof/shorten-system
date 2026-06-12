package security

import (
	"crypto/sha256"
	"encoding/hex"
)

type RefreshTokenHash interface {
	Hash(token string) string
}

type SHA256RefreshTokenHash struct{}

func NewRefreshTokenHash() RefreshTokenHash {
	return &SHA256RefreshTokenHash{}
}

func (h *SHA256RefreshTokenHash) Hash(
	token string,
) string {

	sum := sha256.Sum256(
		[]byte(token),
	)

	return hex.EncodeToString(
		sum[:],
	)
}

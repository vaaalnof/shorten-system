package security

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type VisitorHash interface {
	Hash(
		ipAddress string,
		userAgent string,
	) string
}

type SHA256VisitorHash struct{}

func NewVisitorHash() *SHA256VisitorHash {

	return &SHA256VisitorHash{}
}

func (h *SHA256VisitorHash) Hash(
	ipAddress string,
	userAgent string,
) string {

	key := strings.TrimSpace(
		ipAddress,
	) + "|" + strings.TrimSpace(
		userAgent,
	)

	hash := sha256.Sum256(
		[]byte(key),
	)

	return hex.EncodeToString(
		hash[:],
	)
}

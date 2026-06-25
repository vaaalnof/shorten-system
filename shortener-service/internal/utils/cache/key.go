package cache

import "strings"

const Namespace = "shortener"

type Key string

const (
	URLKeyType Key = "shorturl"
)

func Build(
	key Key,
	parts ...string,
) string {

	base := Namespace +
		":" +
		string(key)

	if len(parts) == 0 {
		return base
	}

	return base +
		":" +
		strings.Join(
			parts,
			":",
		)
}

// ====================================================
// URL
// ====================================================

func URL(
	shortCode string,
) string {

	return Build(
		URLKeyType,
		strings.ToLower(
			strings.TrimSpace(
				shortCode,
			),
		),
	)
}

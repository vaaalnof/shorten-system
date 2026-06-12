package cache

import "strings"

const Namespace = "auth"

type Key string

const (
	SessionKeyType           Key = "session"
	RateLimitLoginKeyType    Key = "rate_limit_login"
	RateLimitRegisterKeyType Key = "rate_limit_register"
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
// SESSION
// ====================================================

func Session(
	sessionID string,
) string {

	return Build(
		SessionKeyType,
		sessionID,
	)
}

// ====================================================
// LOGIN RATE LIMIT
// ====================================================

func RateLimitLogin(
	ip string,
	email string,
) string {

	return Build(
		RateLimitLoginKeyType,
		ip,
		strings.ToLower(
			strings.TrimSpace(
				email,
			),
		),
	)
}

// ====================================================
// REGISTER RATE LIMIT
// ====================================================

func RateLimitRegister(
	ip string,
) string {

	return Build(
		RateLimitRegisterKeyType,
		ip,
	)
}

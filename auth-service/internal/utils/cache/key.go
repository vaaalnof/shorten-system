package cache

import "strings"

const Namespace = "auth"

type Key string

const (
	SessionKeyType                   Key = "session"
	RateLimitLoginKeyType            Key = "rate_limit_login"
	RateLimitRegisterKeyType         Key = "rate_limit_register"
	EmailVerificationKeyType         Key = "email_verification"
	EmailVerificationCooldownKeyType Key = "email_verification_cooldown"
	OAuthStateKeyType                Key = "oauth_state"
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

// ====================================================
// EMAIL VERIFICATION
// ====================================================

func EmailVerification(
	userID string,
) string {

	return Build(
		EmailVerificationKeyType,
		userID,
	)
}

// ====================================================
// EMAIL VERIFICATION COOLDOWN
// ====================================================

func EmailVerificationCooldown(
	userID string,
) string {

	return Build(
		EmailVerificationCooldownKeyType,
		userID,
	)
}

// ====================================================
// OAUTH STATE
// ====================================================

func OAuthState(
	provider string,
	state string,
) string {

	return Build(
		OAuthStateKeyType,
		provider,
		state,
	)
}

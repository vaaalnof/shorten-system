package query

const (
	AddAuthProvider = `
		INSERT INTO user_auth_providers (
			id,
			user_id,
			provider,
			provider_user_id,
			password_hash,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	FindByEmailAndProvider = `
		SELECT
			ap.id,
			ap.user_id,
			ap.provider,
			ap.provider_user_id,
			ap.password_hash,
			ap.created_at
		FROM user_auth_providers ap
		JOIN users u
			ON u.id = ap.user_id
		WHERE u.email = $1
		AND ap.provider = $2
		LIMIT 1
	`

	FindByProviderUserID = `
		SELECT
			id,
			user_id,
			provider,
			provider_user_id,
			password_hash,
			created_at
		FROM user_auth_providers
		WHERE provider = $1
		AND provider_user_id = $2
		LIMIT 1
	`

	FindByUserIDAndProvider = `
		SELECT
			id,
			user_id,
			provider,
			provider_user_id,
			password_hash,
			created_at
		FROM user_auth_providers
		WHERE user_id = $1
		AND provider = $2
		LIMIT 1
	`
)

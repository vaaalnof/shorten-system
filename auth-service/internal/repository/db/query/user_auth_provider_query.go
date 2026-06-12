package query

const (
	UserAuthProviderCreate = `
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

	UserAuthProviderFindLocalByEmail = `
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
		AND ap.provider = 'local'
		LIMIT 1
	`
)

package query

const (
	UserSessionCreate = `
		INSERT INTO user_sessions (
			id,
			user_id,
			refresh_token,
			ip_address,
			user_agent,
			expired_at,
			revoked_at,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	UserSessionFindByID = `
		SELECT
			id,
			user_id,
			refresh_token,
			ip_address,
			user_agent,
			expired_at,
			revoked_at,
			created_at
		FROM user_sessions
		WHERE id = $1
		LIMIT 1
	`

	UserSessionFindValidByID = `
		SELECT
			id,
			user_id,
			refresh_token,
			ip_address,
			user_agent,
			expired_at,
			revoked_at,
			created_at
		FROM user_sessions
		WHERE id = $1
		AND revoked_at IS NULL
		AND expired_at > $2
		LIMIT 1
	`

	UserSessionUpdateRefreshToken = `
		UPDATE user_sessions
		SET refresh_token = $1
		WHERE id = $2
	`

	UserSessionRevoke = `
		UPDATE user_sessions
		SET revoked_at = $1
		WHERE id = $2
		AND revoked_at IS NULL
	`
)

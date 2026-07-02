package query

const (
	AddSession = `
		INSERT INTO user_sessions (
			id,
			user_id,
			refresh_token,
			ip_address,
			user_agent,
			last_seen_at,
			expired_at,
			revoked_at,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	FindSessionByID = `
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

	FindSessionByUserID = `
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
		WHERE user_id = $1
		AND revoked_at IS NULL
		AND expired_at > $2
	`

	UpdateRefreshToken = `
		UPDATE user_sessions
		SET
			refresh_token = $1,
			last_seen_at = $2
		WHERE id = $3
	`

	RevokeByID = `
		UPDATE user_sessions
		SET revoked_at = $1
		WHERE id = $2
		AND revoked_at IS NULL
	`

	RevokeByUserID = `
		UPDATE user_sessions
		SET revoked_at = $1
		WHERE user_id = $2
		AND revoked_at IS NULL
	`
)

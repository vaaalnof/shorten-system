package query

const (
	AddUser = `
		INSERT INTO users (
			id,
			email,
			first_name,
			last_name,
			avatar_url,
			is_active,
			email_verified,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	FindByEmail = `
		SELECT
			id,
			email,
			first_name,
			last_name,
			avatar_url,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	FindByUserID = `
		SELECT
			id,
			email,
			first_name,
			last_name,
			avatar_url,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	UpdateEmailVerified = `
		UPDATE users
		SET
			email_verified = TRUE,
			email_verified_at = $1,
			updated_at = $2
		WHERE id = $3
	`
)

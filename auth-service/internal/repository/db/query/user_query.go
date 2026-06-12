package query

const (
	UserCreate = `
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

	UserFindByEmail = `
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

	UserFindByID = `
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
)

package query

const (

	// =====================================================
	// CREATE
	// =====================================================

	AddURL = `
		INSERT INTO urls (
			id,
			user_id,
			short_code,
			original_url,
			is_active,
			password_hash,
			expired_at,
			deleted_at,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		)
	`

	// =====================================================
	// FIND BY SHORT CODE
	// =====================================================

	FindURLByShortCode = `
		SELECT
			id,
			user_id,
			short_code,
			original_url,
			is_active,
			password_hash,
			expired_at,
			deleted_at,
			created_at,
			updated_at
		FROM urls
		WHERE short_code = $1
		AND is_active = true
		AND (expired_at IS NULL OR expired_at > EXTRACT(EPOCH FROM NOW()))
		LIMIT 1
	`

	// =====================================================
	// FIND BY ID
	// =====================================================

	FindURLByID = `
		SELECT
			id,
			user_id,
			short_code,
			original_url,
			is_active,
			password_hash,
			expired_at,
			deleted_at,
			created_at,
			updated_at
		FROM urls
		WHERE id = $1
		AND deleted_at IS NULL
		LIMIT 1
	`

	// =====================================================
	// UPDATE PASSWORD
	// =====================================================

	UpdateURLPassword = `
		UPDATE urls
		SET
			password_hash = $3,
			updated_at = $4
		WHERE id = $1
		AND user_id = $2
		AND deleted_at IS NULL
	`

	// =====================================================
	// REMOVE PASSWORD
	// =====================================================

	RemoveURLPassword = `
		UPDATE urls
		SET
			password_hash = NULL,
			updated_at = $3
		WHERE id = $1
		AND user_id = $2
		AND deleted_at IS NULL
	`

	// =====================================================
	// COUNT URLS BY USER
	// =====================================================

	CountURLsByUserID = `
		SELECT COUNT(*)
		FROM urls
		WHERE user_id = $1
		AND deleted_at IS NULL
`

	// =====================================================
	// LIST URLS BY USER
	// =====================================================

	ListURLsByUserID = `
		SELECT
			id,
			user_id,
			short_code,
			original_url,
			is_active,
			password_hash,
			expired_at,
			deleted_at,
			created_at,
			updated_at
		FROM urls
		WHERE user_id = $1
		AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3
	`

	// =====================================================
	// DELETE URL
	// =====================================================

	DeleteURL = `
		UPDATE urls
		SET
			deleted_at = $3,
			updated_at = $4
		WHERE id = $1
		AND user_id = $2
		AND deleted_at IS NULL
	`

	// =====================================================
	// UPDATE URL
	// =====================================================

	UpdateURL = `
		UPDATE urls
		SET
			original_url = $3,
			short_code = $4,
			is_active = $5,
			expired_at = $6,
			updated_at = $7
		WHERE id = $1
		AND user_id = $2
		AND deleted_at IS NULL
	`
)

package query

const (
	ReservedAliasExists = `
		SELECT EXISTS (
			SELECT 1
			FROM reserved_aliases
			WHERE LOWER(keyword) = LOWER($1)
		)
	`
)

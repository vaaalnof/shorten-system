package port

import "context"

type ReservedAliasRepository interface {
	Exists(
		ctx context.Context,
		keyword string,
	) (bool, error)
}

package setup

import (
	"context"

	"github.com/leandrohsilveira/tsm/database"
)

func SetupDatabasePool(ctx context.Context) (context.Context, error) {
	pool, err := database.NewPool(ctx)

	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, database.DatabasePoolKey, pool), nil
}

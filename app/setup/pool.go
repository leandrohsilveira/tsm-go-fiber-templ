package setup

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/database"
)

func SetupDatabasePool(ctx context.Context, app *fiber.App) (context.Context, error) {
	pool, err := database.NewPool(ctx)

	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, database.DatabasePoolKey, pool), nil
}

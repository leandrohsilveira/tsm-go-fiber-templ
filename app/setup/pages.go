package setup

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/leandrohsilveira/tsm/database"
	"github.com/leandrohsilveira/tsm/user"
)

func SetupPages(ctx context.Context, app *fiber.App) {
	pool := ctx.Value(database.DatabasePoolKey).(database.DatabasePool)

	app.Static("/public", "./public")

	userService := user.NewService(pool)
	userController := user.NewController(userService)
	app.Mount("/user", user.Pages(userController))
}

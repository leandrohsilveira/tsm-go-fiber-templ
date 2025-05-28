package setup

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/auth"
	"github.com/leandrohsilveira/tsm/database"
	"github.com/leandrohsilveira/tsm/home"
	"github.com/leandrohsilveira/tsm/user"
)

func SetupPages(ctx context.Context, app *fiber.App) {
	pool := ctx.Value(database.DatabasePoolKey).(database.DatabasePool)

	userService := user.NewService(pool)
	authService := auth.NewService(userService)
	userController := user.NewController(userService)
	authController := auth.NewController(authService)

	app.Static("/public", "./public")
	app.Use(auth.AuthMiddleware(authController))
	app.Mount("/", home.Pages())
	app.Mount("/user", user.Pages(userController))
	app.Mount("/login", auth.Pages(authController))
}

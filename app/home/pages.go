package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/guards"
	"github.com/leandrohsilveira/tsm/render"
)

func Pages() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		return render.Html(c, HomePage(info))
	})

	return app
}

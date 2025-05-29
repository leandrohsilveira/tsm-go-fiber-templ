package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func SetupETag(app *fiber.App) {
	app.Use(etag.New(etag.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Response().StatusCode() >= 300
		},
	}))
}

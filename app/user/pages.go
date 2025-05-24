package user

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/leandrohsilveira/tsm/user/templates"
	"github.com/rs/zerolog/log"
)

func Pages(controller UserController) *fiber.App {

	app := fiber.New()

	app.Get("/new", adaptor.HTTPHandler(
		templ.Handler(
			user_templates.SignUpPage(),
		),
	))

	app.Post("/new", func(c *fiber.Ctx) error {
		log.
			Ctx(c.UserContext()).
			Info().
			Msg("Create user form submitted")
		return c.Redirect("/", http.StatusMovedPermanently)
	})

	return app
}

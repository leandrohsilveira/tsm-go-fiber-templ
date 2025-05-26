package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/render"
	"github.com/leandrohsilveira/tsm/user/templates"
	"github.com/rs/zerolog/log"
)

func Pages(controller UserController) *fiber.App {

	app := fiber.New()

	app.Get("/new", func(c *fiber.Ctx) error {
		return render.Html(c, user_templates.SignUpPage(c.Path()))
	})

	app.Post("/new", func(c *fiber.Ctx) error {

		log.
			Ctx(c.UserContext()).
			Info().
			Msg("Create user form submitted")

		return c.Redirect("/", http.StatusMovedPermanently)
	})

	return app
}

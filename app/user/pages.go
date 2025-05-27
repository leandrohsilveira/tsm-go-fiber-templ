package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/render"
	"github.com/rs/zerolog/log"
)

func Pages(controller UserController) *fiber.App {
	app := fiber.New()

	app.Get("/signup", func(c *fiber.Ctx) error {
		return render.Html(c, SignUpPage(c.Path(), nil))
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		_, validationErr, err := controller.Create(c)

		if err != nil {
			return err
		}

		if validationErr != nil {
			return render.Html(c, SignUpPage(c.Path(), validationErr))
		}

		logger := log.Ctx(c.UserContext())
		logger.Info().Msg("Create user form submitted")

		return c.Redirect("/", http.StatusMovedPermanently)
	})

	return app
}

package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/render"
)

func Pages(controller AuthController) *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return render.Html(c, LoginPage(c.Path(), nil, nil))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		result, validationErr, err := controller.Login(c)

		if err == AuthUsernamePasswordIncorrectErr {
			return render.Html(c, LoginPage(c.Path(), nil, err))
		}

		if err != nil {
			return render.Html(c, LoginPage(c.Path(), nil, render.DefaultErr(err, "Authentication failed")))
		}

		if validationErr != nil {
			return render.Html(c, LoginPage(c.Path(), validationErr, nil))
		}

		c.Cookie(&fiber.Cookie{
			Name:     "Authorization",
			Value:    result.Token,
			Path:     "/",
			HTTPOnly: true,
			Secure:   true,
		})

		return c.Redirect("/", http.StatusMovedPermanently)
	})

	return app

}

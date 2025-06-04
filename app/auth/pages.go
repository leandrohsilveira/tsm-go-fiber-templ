package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/guards"
	"github.com/leandrohsilveira/tsm/render"
)

func Pages(controller AuthController) *fiber.App {
	app := fiber.New()

	app.Get("/login", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		return render.Html(c, LoginPage(LoginPageProps{Action: c.Path()}))
	})

	app.Post("/login", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		response, err := controller.Login(c)

		if err == AuthUsernamePasswordIncorrectErr {
			return render.Html(c, LoginPage(LoginPageProps{Action: c.Path(), Err: err}))
		}

		if err != nil {
			return render.Html(c, LoginPage(LoginPageProps{Action: c.Path(), Err: render.DefaultErr(c, err, "Authentication failed")}))
		}

		if response.ValidationErr != nil {
			return render.Html(c, LoginPage(LoginPageProps{
				Action:        c.Path(),
				ValidationErr: response.ValidationErr,
				Value:         response.Payload,
			}))
		}

		c.Cookie(&fiber.Cookie{
			Name:     "Authorization",
			Value:    response.Result.Token,
			Path:     "/",
			HTTPOnly: true,
			Secure:   true,
		})

		return c.Redirect("/", http.StatusSeeOther)
	})

	return app

}

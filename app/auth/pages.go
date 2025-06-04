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
		return render.Html(c, LoginPage(LoginPageProps{Action: c.OriginalURL()}))
	})

	app.Post("/login", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		response, err := controller.Login(c)

		if err == AuthUsernamePasswordIncorrectErr {
			return render.Html(c, LoginPage(LoginPageProps{Action: c.OriginalURL(), Err: err}))
		}

		if err != nil {
			return render.Html(c, LoginPage(LoginPageProps{Action: c.OriginalURL(), Err: render.DefaultErr(c, err, "Authentication failed")}))
		}

		if response.ValidationErr != nil {
			return render.Html(c, LoginPage(LoginPageProps{
				Action:        c.OriginalURL(),
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

		return c.Redirect(c.Query("next", "/"), http.StatusSeeOther)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie("Authorization")
		return c.Redirect("/", http.StatusSeeOther)
	})

	app.Get("/auth/change-password", guards.RegularUserGuard, func(c *fiber.Ctx) error {
		return render.Html(c, ChangeCurrentPasswordPage(ChangeCurrentPasswordPageProps{
			CurrentUserInfo: guards.GetCurrentUser(c),
			Action:          c.Path(),
			BackUrl:         "/",
		}))
	})

	app.Post("/auth/change-password", guards.RegularUserGuard, func(c *fiber.Ctx) error {
		response, err := controller.ChangePassword(c)

		if err != nil {
			return render.Html(c, ChangeCurrentPasswordPage(ChangeCurrentPasswordPageProps{
				CurrentUserInfo: guards.GetCurrentUser(c),
				Action:          c.Path(),
				BackUrl:         "/",
				ValidationErr:   response.ValidationErr,
				Err:             render.DefaultErr(c, err, "Update password failed"),
			}))
		}

		if response.ValidationErr != nil {
			return render.Html(c, ChangeCurrentPasswordPage(ChangeCurrentPasswordPageProps{
				CurrentUserInfo: guards.GetCurrentUser(c),
				Action:          c.Path(),
				BackUrl:         "/",
				ValidationErr:   response.ValidationErr,
			}))
		}

		return c.Redirect("/", http.StatusSeeOther)
	})

	return app

}

package user

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/components/layout"
	"github.com/leandrohsilveira/tsm/guards"
	"github.com/leandrohsilveira/tsm/render"
	"github.com/rs/zerolog/log"
)

func Pages(controller UserController) *fiber.App {
	app := fiber.New()

	app.Get("/signup", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		return render.Html(c, SignUpPage(SignUpPageProps{
			Action:  c.Path(),
			BackUrl: "/",
		}))
	})

	app.Post("/signup", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		response, err := controller.Create(c)

		if err != nil && response.Payload == nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Err:     render.DefaultErr(c, err, "Unable to parse sign-up request payload"),
				BackUrl: "/signup",
			}))
		}

		if err != nil {
			return render.Html(c, SignUpPage(SignUpPageProps{
				Action:  c.Path(),
				BackUrl: "/",
				Err:     render.DefaultErr(c, err, "User sign-up failed"),
				Value:   *response.Payload,
			}))
		}

		if response.ValidationErr != nil {
			return render.Html(c, SignUpPage(SignUpPageProps{
				Action:        c.Path(),
				BackUrl:       "/",
				ValidationErr: response.ValidationErr,
				Value:         *response.Payload,
			}))
		}

		logger := log.Ctx(c.UserContext())
		logger.Info().Msg("User sign-up successful")

		return c.Redirect("/", http.StatusSeeOther)
	})

	app.Get("/manage", guards.AdminUserGuard, func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		result, err := controller.List(c)

		if err != nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Info: info,
				Err:  render.DefaultErr(c, err, "Unable to load users list"),
			}))
		}

		return render.Html(c, UserManagePage(result.Items, info))
	})

	app.Get("/manage/:id", guards.AdminUserGuard, func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		data, err := controller.GetByID(c)

		if err != nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Info:    info,
				Err:     render.DefaultErr(c, err, "Unable to load user data"),
				BackUrl: "../manage",
			}))
		}

		if data == nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Info:    info,
				Err:     errors.New("User not found"),
				BackUrl: "../manage",
			}))
		}

		return render.Html(c, UserManageEditPage(UserManageEditPageProps{
			Value:           data,
			CurrentUserInfo: info,
			Action:          c.Path(),
			BackUrl:         "../manage",
		}))
	})

	app.Post("/manage/:id", guards.AdminUserGuard, func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		response, err := controller.Update(c)

		if err != nil && response.Payload == nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Err:     render.DefaultErr(c, err, "Unable to parse update user request payload"),
				Info:    info,
				BackUrl: c.Path(),
			}))
		}

		if err != nil {
			return render.Html(c, UserManageEditPage(UserManageEditPageProps{
				Err:             err,
				CurrentUserInfo: info,
				Action:          c.Path(),
				BackUrl:         "../manage",
				Value: &UserDisplayDto{
					Name:  response.Payload.Name,
					Email: response.Payload.Email,
					Role:  response.Payload.Role,
				},
			}))
		}

		if response.ValidationErr != nil {
			return render.Html(c, UserManageEditPage(UserManageEditPageProps{
				CurrentUserInfo: info,
				Action:          c.Path(),
				ValidationErr:   response.ValidationErr,
				BackUrl:         "../manage",
				Value: &UserDisplayDto{
					Name:  response.Payload.Name,
					Email: response.Payload.Email,
					Role:  response.Payload.Role,
				},
			}))
		}

		if response.Result == nil {
			return render.Html(c, layout.Error(layout.ErrorProps{
				Err:     errors.New("Unable to update an user that does not exists"),
				Info:    info,
				BackUrl: "../manage",
			}))
		}

		return c.Redirect("../manage", http.StatusSeeOther)
	})

	return app
}

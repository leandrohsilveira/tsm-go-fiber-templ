package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/guards"
	"github.com/leandrohsilveira/tsm/render"
	"github.com/rs/zerolog/log"
)

func Pages(controller UserController) *fiber.App {
	app := fiber.New()

	app.Get("/signup", guards.AnonymousGuard, func(c *fiber.Ctx) error {
		return render.Html(c, SignUpPage(c.Path(), nil))
	})

	app.Post("/signup", guards.AnonymousGuard, func(c *fiber.Ctx) error {
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

	app.Get("/manage", guards.AdminUserGuard, func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		result, err := controller.List(c)

		if err != nil {
			return err
		}

		return render.Html(c, UserManagePage(result.Items, info))
	})

	app.Get("/manage/:id", guards.AdminUserGuard, func(c *fiber.Ctx) error {
		info := guards.GetCurrentUser(c)

		data, err := controller.GetByID(c)

		if err != nil {
			return err
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

		data, validationErr, err := controller.Update(c)

		if err != nil {
			return err
		}

		if validationErr != nil {
			return render.Html(c, UserManageEditPage(UserManageEditPageProps{
				Value: &UserDisplayDto{
					Name:  validationErr.Data["Name"].(string),
					Email: validationErr.Data["Email"].(string),
					Role:  validationErr.Data["Role"].(dao.UserRole),
				},
				CurrentUserInfo: info,
				Action:          c.Path(),
				ValidationErr:   validationErr,
				BackUrl:         "../manage",
			}))
		}

		if data == nil {
			return c.Redirect("../manage", http.StatusSeeOther)
		}

		return c.Redirect("../manage", http.StatusSeeOther)
	})

	return app
}

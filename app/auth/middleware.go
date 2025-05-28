package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/guards"
)

func AuthMiddleware(controller AuthController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		info, err := controller.GetCurrentUser(c)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		if info != nil {
			guards.SetCurrentUser(c, guards.CurrentUserDto{
				ID:    info.User.ID,
				Name:  info.User.Name,
				Email: info.User.Email,
				Role:  info.User.Role,
			})
		}

		return c.Next()
	}
}

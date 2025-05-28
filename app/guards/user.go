package guards

import (
	"github.com/gofiber/fiber/v2"
)

func RegularUserGuard(c *fiber.Ctx) error {
	info := GetCurrentUser(c)

	if info == nil {
		return redirectLogin(c)
	}

	return c.Next()
}

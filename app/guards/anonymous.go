package guards

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AnonymousGuard(c *fiber.Ctx) error {
	info := GetCurrentUser(c)

	if info == nil {
		return c.Next()
	}

	return c.Redirect("/", http.StatusSeeOther)
}

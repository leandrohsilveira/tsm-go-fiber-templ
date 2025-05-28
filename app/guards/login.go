package guards

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func redirectLogin(c *fiber.Ctx) error {
	return c.Redirect("/login", http.StatusTemporaryRedirect)
}

package guards

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func redirectLogin(c *fiber.Ctx) error {
	path := c.OriginalURL()
	return c.Redirect(fmt.Sprintf("/login?next=%s", path), http.StatusSeeOther)
}

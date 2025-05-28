package guards

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

var authCurrentUserKey struct{}

func SetCurrentUser(c *fiber.Ctx, data CurrentUserDto) {
	c.SetUserContext(
		context.WithValue(c.UserContext(), authCurrentUserKey, data),
	)
}

func GetCurrentUser(c *fiber.Ctx) *CurrentUserDto {
	data, ok := c.UserContext().Value(authCurrentUserKey).(CurrentUserDto)
	if !ok {
		return nil
	}
	return &data
}

package setup

import (
	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: "TSM",
		ETag:    true,
	})
}

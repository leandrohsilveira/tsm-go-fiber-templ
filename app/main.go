package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/setup"
)

func main() {
	app := setup.SetupApp()

	setup.SetupLogger(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}

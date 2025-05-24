package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/setup"
)

func main() {
	ctx := context.Background()
	app := setup.SetupApp()

	setup.SetupLogger(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	ctx, err := setup.SetupDatabasePool(ctx, app)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(":3000"))
}

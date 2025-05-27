package render

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Html(ctx *fiber.Ctx, component templ.Component) error {
	ctx.Set("Content-Type", "text/html")
	return component.Render(ctx.Context(), ctx.Response().BodyWriter())
}

func DefaultErr(err error, message string) error {
	if message == "" {
		message = "Internal server error"
	}
	return errors.Join(errors.New(message), err)
}

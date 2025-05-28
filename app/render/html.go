package render

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Html(ctx *fiber.Ctx, component templ.Component) error {
	ctx.Set("Content-Type", "text/html")
	return component.Render(ctx.Context(), ctx.Response().BodyWriter())
}

func DefaultErr(ctx *fiber.Ctx, err error, message string) error {
	if message == "" {
		message = "Internal server error"
	}
	log.Ctx(ctx.UserContext()).Error().Stack().Err(err).Msg(message)
	return errors.New(message)
}

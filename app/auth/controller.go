package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/util"
)

type AuthController interface {
	Login(*fiber.Ctx) (*AuthLoginResultDto, *util.ValidationErr[AuthLoginPayloadDto], error)
}

type authController struct {
	authService AuthService
}

func (self *authController) Login(c *fiber.Ctx) (*AuthLoginResultDto, *util.ValidationErr[AuthLoginPayloadDto], error) {
	payload := new(AuthLoginPayloadDto)

	if err := c.BodyParser(payload); err != nil {
		return nil, nil, fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	validationErr, err := util.Validate(payload)

	if err != nil {
		return nil, nil, err
	}

	if validationErr != nil {
		return nil, validationErr, nil
	}

	result, err := self.authService.Login(c.Context(), *payload)

	if err == AuthUsernamePasswordIncorrectErr {
		return nil, nil, fiber.NewError(http.StatusForbidden, err.Error())
	}

	if err != nil {
		return nil, nil, err
	}

	return result, nil, nil
}

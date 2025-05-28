package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/util"
)

type AuthController interface {
	Login(*fiber.Ctx) (*AuthLoginResultDto, *util.ValidationErr[AuthLoginPayloadDto], error)
	GetCurrentUser(*fiber.Ctx) (*AuthCurrentUserInfoDto, error)
}

type authController struct {
	authService AuthService
}

func NewController(authService AuthService) AuthController {
	return &authController{authService: authService}
}

func (self *authController) Login(c *fiber.Ctx) (*AuthLoginResultDto, *util.ValidationErr[AuthLoginPayloadDto], error) {
	payload := new(AuthLoginPayloadDto)

	if err := c.BodyParser(payload); err != nil {
		return nil, nil, err
	}

	validationErr, err := util.Validate(payload)

	if err != nil {
		return nil, nil, err
	}

	if validationErr != nil {
		return nil, validationErr, nil
	}

	result, err := self.authService.Login(c.Context(), *payload)

	if err != nil {
		return nil, nil, err
	}

	return result, nil, nil
}

func (self *authController) GetCurrentUser(c *fiber.Ctx) (*AuthCurrentUserInfoDto, error) {
	token := c.Get("authorization")

	if token == "" {
		token = c.Cookies("Authorization")
	}

	if token == "" {
		return nil, nil
	}

	return self.authService.GetCurrentUserInfo(c.Context(), token)
}

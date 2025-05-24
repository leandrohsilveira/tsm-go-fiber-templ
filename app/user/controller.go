package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Create(*fiber.Ctx) (*UserDisplayDto, error)
}

type userController struct {
	userService UserService
}

func NewController(userService UserService) UserController {
	return &userController{userService: userService}
}

func (controller *userController) Create(c *fiber.Ctx) (*UserDisplayDto, error) {
	payload := new(UserCreateDto)

	if err := c.BodyParser(payload); err != nil {
		return nil, fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	data, err := controller.userService.Create(c.UserContext(), *payload)

	if err != nil {
		return nil, err
	}

	dto := &UserDisplayDto{
		ID:    data.ID.String(),
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}

	return dto, nil
}

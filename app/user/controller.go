package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/util"
)

type UserController interface {
	Create(*fiber.Ctx) (*UserDisplayDto, *util.ValidationErr[UserCreateDto], error)
	List(*fiber.Ctx) (*util.PageResult[UserDisplayDto], error)
}

type userController struct {
	userService UserService
}

func NewController(userService UserService) UserController {
	return &userController{userService: userService}
}

func (controller *userController) Create(c *fiber.Ctx) (*UserDisplayDto, *util.ValidationErr[UserCreateDto], error) {
	payload := new(UserCreateDto)

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

	data, err := controller.userService.Create(c.UserContext(), *payload)

	if err != nil {
		return nil, nil, err
	}

	dto := &UserDisplayDto{
		ID:    data.ID.String(),
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}

	return dto, nil, nil
}

func (controller *userController) List(c *fiber.Ctx) (*util.PageResult[UserDisplayDto], error) {
	params := new(util.PageParams)

	if err := c.QueryParser(params); err != nil {
		return nil, fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	params.SetDefaults()

	result, err := controller.userService.GetList(c.UserContext(), *params)

	if err != nil {
		return nil, err
	}

	dto := util.MapResult(*result, func(item dao.User) UserDisplayDto {
		return UserDisplayDto{
			ID:    item.ID.String(),
			Name:  item.Name,
			Email: item.Email,
			Role:  item.Role,
		}
	})

	return &dto, nil

}

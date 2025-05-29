package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/util"
)

type UserController interface {
	Create(*fiber.Ctx) (*UserDisplayDto, *util.ValidationErr, error)
	List(*fiber.Ctx) (*util.PageResult[UserDisplayDto], error)
	GetByID(*fiber.Ctx) (*UserDisplayDto, error)
}

type userController struct {
	userService UserService
}

func NewController(userService UserService) UserController {
	return &userController{userService: userService}
}

func (controller *userController) Create(c *fiber.Ctx) (*UserDisplayDto, *util.ValidationErr, error) {
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

func (controller *userController) GetByID(c *fiber.Ctx) (*UserDisplayDto, error) {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return nil, fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	data, err := controller.userService.GetByID(c.UserContext(), id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	result := &UserDisplayDto{
		ID:    data.ID.String(),
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}

	return result, nil
}

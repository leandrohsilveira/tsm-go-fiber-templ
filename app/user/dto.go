package user

import (
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/util"
)

type UserCreatePayloadDto struct {
	Name            string       `json:"name" form:"name" validate:"required"`
	Email           string       `json:"email" form:"email" validate:"required,email"`
	Password        string       `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string       `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
	Role            dao.UserRole `json:"role" form:"role"`
}

type UserCreateResponseDto struct {
	Payload       *UserCreatePayloadDto
	ValidationErr *util.ValidationErr
	Result        *UserDisplayDto
}

type UserManageEditPayloadDto struct {
	Name  string       `json:"name" form:"name" validate:"required"`
	Email string       `json:"email" form:"email" validate:"required,email"`
	Role  dao.UserRole `json:"role" form:"role" validate:"required"`
}

type UserManageEditResponseDto struct {
	Payload       *UserManageEditPayloadDto
	ValidationErr *util.ValidationErr
	Result        *UserDisplayDto
}

type UserChangePasswordDto struct {
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type UserDisplayDto struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Role  dao.UserRole `json:"role"`
}

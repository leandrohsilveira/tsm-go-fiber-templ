package user

import "github.com/leandrohsilveira/tsm/dao"

type UserCreateDto struct {
	Name            string       `json:"name" form:"name" validate:"required"`
	Email           string       `json:"email" form:"email" validate:"required,email"`
	Password        string       `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string       `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
	Role            dao.UserRole `json:"role" form:"role"`
}

type UserDisplayDto struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Role  dao.UserRole `json:"role"`
}

package auth

import (
	"github.com/google/uuid"
	"github.com/leandrohsilveira/tsm/user"
	"github.com/leandrohsilveira/tsm/util"
)

type AuthLoginPayloadDto struct {
	Username string `json:"username" form:"username" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type AuthLoginResultDto struct {
	Token        string              `json:"token"`
	RefreshToken string              `json:"refresh_token"`
	User         user.UserDisplayDto `json:"user"`
}

type AuthLoginResponseDto struct {
	Payload       AuthLoginPayloadDto
	ValidationErr *util.ValidationErr
	Result        AuthLoginResultDto
}

type AuthChangeCurrentPasswordDto struct {
	CurrentPassword string `json:"current_password" form:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type AuthChangeCurrentPasswordResponseDto struct {
	Payload       AuthChangeCurrentPasswordDto
	ValidationErr *util.ValidationErr
}

type AuthCurrentUserInfoDto struct {
	User user.UserDisplayDto `json:"user"`
}

type AuthTokenPayload struct {
	UserID uuid.UUID
}

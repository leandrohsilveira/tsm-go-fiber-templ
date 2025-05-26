package auth

import (
	"github.com/google/uuid"
	"github.com/leandrohsilveira/tsm/user"
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

type AuthCurrentUserInfoDto struct {
	User user.UserDisplayDto `json:"user"`
}

type AuthTokenPayload struct {
	UserID uuid.UUID
}

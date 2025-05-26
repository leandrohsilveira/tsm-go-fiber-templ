package auth

import (
	"context"
	"errors"

	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/user"
)

var AuthUsernamePasswordIncorrectErr = errors.New("Incorrect username or password")

type AuthService interface {
	Login(context.Context, AuthLoginPayloadDto) (*AuthLoginResultDto, error)
	GetCurrentUserInfo(context.Context, string) (*AuthCurrentUserInfoDto, error)
}

func NewService(userService user.UserService) AuthService {
	return &authService{userService: userService}
}

type authService struct {
	userService user.UserService
}

func (service *authService) Login(ctx context.Context, payload AuthLoginPayloadDto) (*AuthLoginResultDto, error) {
	data, err := service.userService.GetByEmail(ctx, payload.Username)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, AuthUsernamePasswordIncorrectErr
	}

	match, err := service.checkPassword(payload.Password, data.Password.String)

	if err != nil {
		return nil, err
	}

	if !match {
		return nil, AuthUsernamePasswordIncorrectErr
	}

	token, err := service.generateToken(data)

	if err != nil {
		return nil, err
	}

	result := &AuthLoginResultDto{
		Token: token,
		// TODO: generate refresh token
		User: user.UserDisplayDto{
			ID:    data.ID.String(),
			Name:  data.Name,
			Email: data.Email,
			Role:  data.Role,
		},
	}

	return result, nil
}

func (_ *authService) GetCurrentUserInfo(context.Context, string) (*AuthCurrentUserInfoDto, error) {
	// TODO: add get current user info implementation
	return nil, nil
}

func (_ *authService) checkPassword(input, hash string) (bool, error) {
	// TODO: use password hasher to check password
	return input == hash, nil
}

func (_ *authService) generateToken(user *dao.User) (string, error) {
	// TODO: generate a JWT token
	return user.ID.String(), nil
}

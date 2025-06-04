package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/user"
)

var AuthUsernamePasswordIncorrectErr = errors.New("Incorrect username or password")
var AuthIncorrectCurrentPasswordErr = errors.New("Incorrect current password")

type AuthService interface {
	Login(context.Context, AuthLoginPayloadDto) (*AuthLoginResultDto, error)
	GetCurrentUserInfo(context.Context, string) (*AuthCurrentUserInfoDto, error)
	ChangePassword(context.Context, uuid.UUID, AuthChangeCurrentPasswordDto) (*AuthCurrentUserInfoDto, error)
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

func (self *authService) GetCurrentUserInfo(ctx context.Context, token string) (*AuthCurrentUserInfoDto, error) {
	payload, err := self.validateAndReadTokenPayload(token)

	if err != nil {
		return nil, err
	}

	data, err := self.userService.GetByID(ctx, payload.UserID)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	dto := user.UserDisplayDto{
		ID:    data.ID.String(),
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}

	return &AuthCurrentUserInfoDto{User: dto}, nil
}

func (self *authService) ChangePassword(ctx context.Context, id uuid.UUID, payload AuthChangeCurrentPasswordDto) (*AuthCurrentUserInfoDto, error) {
	data, err := self.userService.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	match, err := self.checkPassword(payload.CurrentPassword, data.Password.String)

	if err != nil {
		return nil, err
	}

	if !match {
		return nil, AuthIncorrectCurrentPasswordErr
	}

	data, err = self.userService.ChangePassword(ctx, id, user.UserChangePasswordDto{
		NewPassword: payload.NewPassword,
	})

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	result := &AuthCurrentUserInfoDto{
		User: user.UserDisplayDto{
			ID:    data.ID.String(),
			Name:  data.Name,
			Email: data.Email,
			Role:  data.Role,
		},
	}

	return result, nil
}

func (_ *authService) checkPassword(input, hash string) (bool, error) {
	// TODO: use password hasher to check password
	return input == hash, nil
}

func (_ *authService) generateToken(user *dao.User) (string, error) {
	// TODO: generate a JWT token
	return user.ID.String(), nil
}

func (_ *authService) validateAndReadTokenPayload(token string) (*AuthTokenPayload, error) {
	userId, err := uuid.Parse(token)

	if err != nil {
		return nil, err
	}

	return &AuthTokenPayload{UserID: userId}, nil
}

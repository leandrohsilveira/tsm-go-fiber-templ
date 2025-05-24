package user

import (
	"context"

	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/database"
)

type UserService interface {
	Create(context.Context, UserCreateDto) (*dao.User, error)
	GetByEmail(context.Context, string) (*dao.User, error)
}

var UserServiceKey struct{}

type userService struct {
	pool database.DatabasePool
}

func NewService(pool database.DatabasePool) UserService {
	return &userService{pool: pool}
}

func (self *userService) Create(ctx context.Context, payload UserCreateDto) (*dao.User, error) {
	queries, release, err := self.pool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer release()

	data, err := queries.CreateUser(ctx, dao.CreateUserParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: self.pool.Text(payload.Password),
		Role:     payload.Role,
	})

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (self *userService) GetByEmail(ctx context.Context, email string) (*dao.User, error) {
	queries, release, err := self.pool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer release()

	data, err := queries.GetUserByEmail(ctx, email)

	if err == database.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}

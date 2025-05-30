package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/leandrohsilveira/tsm/database"
	"github.com/leandrohsilveira/tsm/util"
)

type UserService interface {
	Create(context.Context, UserCreateDto) (*dao.User, error)
	Update(context.Context, uuid.UUID, UserManageEditDto) (*dao.User, error)
	GetByEmail(context.Context, string) (*dao.User, error)
	GetByID(context.Context, uuid.UUID) (*dao.User, error)
	GetList(context.Context, util.PageParams) (*util.PageResult[dao.User], error)
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

	role := payload.Role
	if role == "" {
		role = dao.UserRoleRegularUser
	}

	data, err := queries.CreateUser(ctx, dao.CreateUserParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: self.pool.Text(payload.Password),
		Role:     role,
	})

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (self *userService) Update(ctx context.Context, id uuid.UUID, payload UserManageEditDto) (*dao.User, error) {
	queries, release, err := self.pool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer release()

	data, err := queries.UpdateUser(ctx, dao.UpdateUserParams{
		ID:    id,
		Name:  payload.Name,
		Email: payload.Email,
		Role:  payload.Role,
	})

	if err == database.ErrNoRows {
		return nil, nil
	}

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

func (self *userService) GetByID(ctx context.Context, id uuid.UUID) (*dao.User, error) {
	queries, release, err := self.pool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer release()

	data, err := queries.GetUserById(ctx, id)

	if err == database.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (self *userService) GetList(ctx context.Context, params util.PageParams) (*util.PageResult[dao.User], error) {
	queries, release, err := self.pool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer release()

	data, err := queries.ListUsers(ctx, dao.ListUsersParams{
		Limit:  params.Limit,
		Offset: params.Offset(),
	})

	if err != nil {
		return nil, err
	}

	count, err := queries.CountUsers(ctx)

	if err != nil {
		return nil, err
	}

	result := util.NewPageResult(data, count, params.Limit)

	return &result, nil
}

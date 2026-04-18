package service

import (
	"context"
	"errors"

	"github.com/mickey-mickser/mini-bank/internal/domain"
	"github.com/mickey-mickser/mini-bank/internal/repository/postgres"
)

type UserService interface {
	CreateUser(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
}
type userService struct {
	userRepo postgres.UserRepo
}

func NewUserService(userRepo postgres.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
func (u *userService) CreateUser(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	if len(input.Password) < 6 {
		return CreateUserOutput{}, errors.New("there are no salmons in the password")
	} // TODO hash password
	user := &domain.User{
		Name:         input.Name,
		Login:        input.Login,
		PasswordHash: input.Password, //TODO password hash
	}
	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return CreateUserOutput{}, err
	}

	return CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil
}

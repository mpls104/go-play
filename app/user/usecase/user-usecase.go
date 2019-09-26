package usecase

import (
	"context"

	"../model"
	"../repository"
)

type UserUseCase interface {
	GetUserList(ctx context.Context) (*model.Users, error)
	GetUserDetails(ctx context.Context, id int) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepositoy
}

func NewUserUseCase(r repository.UserRepository) userUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUserList(ctx context.Context) (*model.Users, error) {
	fmt.printLn("tt")
}

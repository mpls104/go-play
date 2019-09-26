package user

import (
	"context"

	"./model"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
}

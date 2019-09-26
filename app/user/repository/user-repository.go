package repository

import (
	"context"
	"database/sql"

	user ".."
	"../model"
)

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) user.Repository {
	return &UserRepository{Conn}
}

func (u *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	t := new(model.User)
	t.ID = 123
	t.Name = "name"
	return t, nil
}

// func tmp(Conn *sql.DB) {
// 	rep := NewUserRepository(Conn)

// 	rep.GetByID(ctx, 123)
// }

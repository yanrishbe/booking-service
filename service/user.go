package service

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
)

type User struct {
	db UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{db: repository}
}

func (us User) Create(ctx context.Context, user model.User) (string, error) {
	return us.db.CreateUser(ctx, user)
}

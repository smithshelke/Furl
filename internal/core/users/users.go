package users

import (
	"context"

	"github.com/smithshelke/flur/internal/core/domain"
	"github.com/smithshelke/flur/internal/ports/db"
)

type UserAPI struct {
	UserRepo db.Users
}

func NewUserAPI(userRepo *db.Users) *UserAPI {
	return &UserAPI{
		UserRepo: *userRepo,
	}
}

func (u *UserAPI) Create(ctx context.Context, users []*domain.User) ([]*domain.User, error) {
	user, err := u.UserRepo.Create(ctx, users)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserAPI) List(ctx context.Context) ([]*domain.User, error) {
	users, err := u.UserRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

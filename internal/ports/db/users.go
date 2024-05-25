package db

import (
	"context"

	"github.com/smithshelke/flur/internal/core/domain"
)

type Users interface {
	Create(ctx context.Context, user []*domain.User) ([]*domain.User, error)
	GetUserByID(ID string) (*domain.User, error)
	List(ctx context.Context) ([]*domain.User, error)
}

package db

import (
	"context"

	"github.com/smithshelke/flur/internal/core/domain"
)

type Blocks interface {
	Create(context.Context, []*domain.Block) ([]*domain.Block, error)
}

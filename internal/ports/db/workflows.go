package db

import (
	"context"

	"github.com/smithshelke/flur/internal/core/domain"
)

type Workflows interface {
	Create(ctx context.Context, workflow *domain.Workflow) (*domain.Workflow, error)
	AssignUser(ctx context.Context, userID string, workflowIDs []string) error
	CreateConnections(ctx context.Context, connections []domain.Connection) error
	GetByUserID(userID string) ([]*domain.Workflow, error)
}

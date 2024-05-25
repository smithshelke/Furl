package repos

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/smithshelke/flur/internal/core/domain"
)

type WorkflowsRepo struct {
	Driver neo4j.DriverWithContext
}

func NewWorkflowsRepo(driver neo4j.DriverWithContext) *WorkflowsRepo {
	return &WorkflowsRepo{
		Driver: driver,
	}
}

func (r *WorkflowsRepo) Create(ctx context.Context, workflow *domain.Workflow) (*domain.Workflow, error) {
	createdWorkflow := &domain.Workflow{}
	result, err := neo4j.ExecuteQuery(ctx, r.Driver,
		"MERGE (w: Workflow {name: $name, created_by: $created_by}) RETURN w",
		map[string]any{
			"name":       workflow.Name,
			"created_by": workflow.CreatedBy,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		key := "w"
		workflow, ok := record.Get(key)
		if !ok {
			return nil, fmt.Errorf("%s not present in return statement of query", key)
		}
		workflowNode := workflow.(neo4j.Node)
		createdWorkflow.Name = workflowNode.GetProperties()["name"].(string)
		createdWorkflow.CreatedBy = workflowNode.GetProperties()["created_by"].(string)
		createdWorkflow.ID = workflowNode.GetElementId()
	}
	return createdWorkflow, nil
}

func (r *WorkflowsRepo) AssignUser(ctx context.Context, userID string, workflowIDs []string) error {
	query := `
				MATCH (w:Workflow) 
				WHERE elementId(w) IN $workflow_ids
				MATCH (u:User)
				WHERE elementId(u) = $user_id
				MERGE (u)-[r:OWNS]->(w)
				RETURN u,r,w
			`
	args := map[string]any{
		"user_id":      userID,
		"workflow_ids": workflowIDs,
	}
	_, err := neo4j.ExecuteQuery(ctx, r.Driver, query, args, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	return err
}

func (r *WorkflowsRepo) CreateConnections(ctx context.Context, connections []domain.Connection) error {
	// todo: use txn and perform operation atomically
	for i := range connections {
		connType, err := getConnectionType(connections[i].Type)
		if err != nil {
			return err
		}
		query := fmt.Sprintf(`
				MATCH (f)
				WHERE elementId(f) = $from
				MATCH (t)
				WHERE elementId(t) = $to
				MERGE (f)-[r:%s]->(t)
				RETURN f,r,t
		`, connType)
		args := map[string]any{
			"from":      connections[i].From.ID,
			"to":        connections[i].To.ID,
			"conn_type": connType,
		}
		_, err = neo4j.ExecuteQuery(ctx, r.Driver, query, args, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorkflowsRepo) GetByUserID(userID string) ([]*domain.Workflow, error) {
	return nil, nil

}

package graphs

import (
	"context"
	"log"

	"github.com/smithshelke/flur/internal/core/domain"
	"github.com/smithshelke/flur/internal/core/graphs/executors/blocks/network"
	"github.com/smithshelke/flur/internal/ports/db"
)

type GraphsAPI struct {
	BlocksRepo   db.Blocks
	WorkflowRepo db.Workflows
}

func NewGraphsAPI(workflowRepo *db.Workflows, blocksRepo *db.Blocks) *GraphsAPI {
	return &GraphsAPI{
		WorkflowRepo: *workflowRepo,
		BlocksRepo:   *blocksRepo,
	}
}

func (g *GraphsAPI) CreateBlocks(ctx context.Context, blocks []*domain.Block) ([]*domain.Block, error) {
	blocks, err := g.BlocksRepo.Create(ctx, blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func (g *GraphsAPI) CreateWorkflow(ctx context.Context, workflow *domain.Workflow) (*domain.Workflow, error) {
	workflow, err := g.WorkflowRepo.Create(ctx, workflow)
	if err != nil {
		return nil, err
	}

	return workflow, nil
}

func (g *GraphsAPI) AssignUserToWorkflows(ctx context.Context, userID string, workflowIDs []string) error {
	err := g.WorkflowRepo.AssignUser(ctx, userID, workflowIDs)
	return err
}

func (g *GraphsAPI) CreateConnections(ctx context.Context, connections []domain.Connection) error {
	err := g.WorkflowRepo.CreateConnections(ctx, connections)
	if err != nil {
		return err
	}
	return nil
}

func (g *GraphsAPI) ExecuteBlock(ctx context.Context, blockID string, response *any) error {
	outputHandler := func(output any) {
		networkOutput := output.(network.NetworkBlockOutput)
		log.Printf("Output: %d \n", networkOutput.StatusCode)
		*response = networkOutput
	}
	blockIO, _ := network.NewNetworkBlockExecutorIOBuilder().
		AddMethod("GET").
		AddHost("https://www.google.com").
		AddOutputHandler(outputHandler).
		Build()
	config := network.NewNetworkBlockExecutorConfig()
	executor := network.NewNetworkBlockExecutorBuilder(ctx).
		WithIO(blockIO).
		WithConfig(config).
		Build()
	return executor.Run()

}

func (g *GraphsAPI) GetWorkflowByID(ID string) domain.Workflow {
	return domain.Workflow{}
}

func (g *GraphsAPI) GetByUserID(userID string) domain.Workflow {
	return domain.Workflow{}
}

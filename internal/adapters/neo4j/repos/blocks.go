package repos

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/smithshelke/flur/internal/core/domain"
)

type BlocksRepo struct {
	Driver neo4j.DriverWithContext
}

func NewBlocksRepo(driver neo4j.DriverWithContext) *BlocksRepo {
	return &BlocksRepo{
		Driver: driver,
	}
}

func (r *BlocksRepo) Create(ctx context.Context, blocks []*domain.Block) ([]*domain.Block, error) {
	createdBlocks := []*domain.Block{}
	for _, block := range blocks {
		result, err := neo4j.ExecuteQuery(ctx, r.Driver,
			"MERGE (b:Block {name: $name, type: $type, created_by: $created_by}) RETURN b",
			map[string]any{
				"name":       block.Name,
				"type":       block.Type,
				"created_by": block.CreatedBy,
			}, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			log.Printf("Failed to create block: %+v\n", err)
			return nil, err
		}
		for _, record := range result.Records {
			key := "b"
			block, ok := record.Get(key)
			if !ok {
				return nil, fmt.Errorf("%s not present in return statement of query", key)
			}
			blockNode, _ := block.(neo4j.Node)
			createdBlocks = append(createdBlocks, &domain.Block{
				ID:        blockNode.ElementId,
				Name:      blockNode.GetProperties()["name"].(string),
				Type:      blockNode.GetProperties()["type"].(string),
				CreatedBy: blockNode.GetProperties()["created_by"].(string),
			})
		}
	}
	return createdBlocks, nil
}

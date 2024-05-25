package blocks

import (
	"github.com/smithshelke/flur/internal/core/domain"
)

func CreateBlockRequestToBlocks(r *CreateBlockRequest) []*domain.Block {
	blocks := []*domain.Block{}
	for _, b := range r.Blocks {
		blocks = append(blocks, &domain.Block{Name: b.Name, Type: b.Type, CreatedBy: b.CreatedBy})
	}
	return blocks
}

func BlocksToCreateBlocksResponse(blocks []*domain.Block) *CreateBlockResponse {
	response := &CreateBlockResponse{}
	for _, b := range blocks {
		response.Blocks = append(response.Blocks, BlockResponse{Name: b.Name, Type: b.Type, CreatedBy: b.CreatedBy})
	}
	return response
}

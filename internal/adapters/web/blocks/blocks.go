package blocks

import (
	"encoding/json"
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/web/common"
	"github.com/smithshelke/flur/internal/core/graphs"
)

type BlocksController struct {
	graphsAPI *graphs.GraphsAPI
}

func NewBlocksController(graphsAPI *graphs.GraphsAPI) *BlocksController {
	return &BlocksController{
		graphsAPI: graphsAPI,
	}
}

func (b *BlocksController) Create(w http.ResponseWriter, r *http.Request) {
	request := &CreateBlockRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}
	blocks, err := b.graphsAPI.CreateBlocks(r.Context(), CreateBlockRequestToBlocks(request))
	if err != nil {
		common.WriteError(w, err, common.SOMETHING_WENT_WRONG)
		return
	}

	response := BlocksToCreateBlocksResponse(blocks)
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (b *BlocksController) Execute(w http.ResponseWriter, r *http.Request) {
	var response any
	err := b.graphsAPI.ExecuteBlock(r.Context(), "123", &response)
	if err != nil {
		common.WriteError(w, err, common.SOMETHING_WENT_WRONG)
		return
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

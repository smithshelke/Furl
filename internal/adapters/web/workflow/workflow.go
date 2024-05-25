package workflow

import (
	"encoding/json"
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/web/common"
	"github.com/smithshelke/flur/internal/core/graphs"
)

type WorkflowController struct {
	GraphsAPI *graphs.GraphsAPI
}

func NewWorkflowController(graphs *graphs.GraphsAPI) *WorkflowController {
	return &WorkflowController{
		GraphsAPI: graphs,
	}
}

func (c *WorkflowController) Create(w http.ResponseWriter, r *http.Request) {
	request := &CreateWorkflowRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}
	workflow, err := c.GraphsAPI.CreateWorkflow(r.Context(), CreateWorkflowRequestToWorkflow(request))
	if err != nil {
		common.WriteError(w, err, common.SOMETHING_WENT_WRONG)
	}
	response, err := json.Marshal(WorkflowToCreateWorkflowResponse(workflow))
	if err != nil {
		common.WriteError(w, err, common.SOMETHING_WENT_WRONG)
	}

	w.Write(response)
	return
}

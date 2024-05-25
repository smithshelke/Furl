package workflow

import "github.com/smithshelke/flur/internal/core/domain"

func CreateWorkflowRequestToWorkflow(request *CreateWorkflowRequest) *domain.Workflow {
	return &domain.Workflow{
		Name:      request.Name,
		CreatedBy: request.CreatedBy,
	}
}

func WorkflowToCreateWorkflowResponse(workflow *domain.Workflow) *CreateWorkflowResponse {
	return &CreateWorkflowResponse{
		ID:        workflow.ID,
		Name:      workflow.Name,
		CreatedBy: workflow.CreatedBy,
	}
}

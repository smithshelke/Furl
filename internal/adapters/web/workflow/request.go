package workflow

type CreateWorkflowRequest struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

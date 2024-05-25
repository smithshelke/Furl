package domain

type Workflow struct {
	ID   string
	Name string
	// CreatedBy is the userID that creates the workflow
	CreatedBy string
}

type WorkflowConfig struct {
	WorflowID            string
	DefaultBlockVariants []DefaultBlockVariants
	Variables            map[string]string
}

type DefaultBlockVariants struct {
	BlockID        string
	BlockVariantID string
}

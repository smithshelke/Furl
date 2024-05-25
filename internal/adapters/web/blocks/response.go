package blocks

type CreateBlockResponse struct {
	Blocks []BlockResponse `json:"blocks"`
}

type BlockResponse struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedBy string `json:"created_by"`
}

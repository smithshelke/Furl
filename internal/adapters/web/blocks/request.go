package blocks

type CreateBlockRequest struct {
	Blocks []BlockRequest `json:"blocks"`
}

type BlockRequest struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedBy string `json:"created_by"`
}

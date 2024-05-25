package connection

type Node struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Connection struct {
	From Node   `json:"from"`
	To   Node   `json:"to"`
	Type string `json:"type"`
}

type CreateConnectionRequest struct {
	Connections []Connection `json:"connections"`
}

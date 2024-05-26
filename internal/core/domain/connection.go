package domain

type ConnectionType int

const (
	TRIGGERS ConnectionType = iota
	STARTS_WITH
	ENDS_WITH
	UNKNOWN_CONNECTION_TYPE
)

type NodeType int

const (
	BLOCK NodeType = iota
	WORKFLOW
	UNKNOWN_NODE_TYPE
)

type Node struct {
	ID   string
	Type NodeType
}

type Connection struct {
	ID   string
	From *Node
	To   *Node
	Type ConnectionType
}

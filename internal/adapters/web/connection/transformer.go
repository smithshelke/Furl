package connection

import (
	"fmt"

	"github.com/smithshelke/flur/internal/core/domain"
)

func CreateConnectionRequestToConnections(request *CreateConnectionRequest) ([]domain.Connection, error) {
	connections := []domain.Connection{}
	for i := range request.Connections {
		conn := request.Connections[i]

		connType, err := getConnectionType(conn.Type)
		if err != nil {
			return nil, err
		}

		fromType, err := getNodeType(conn.From.Type)
		if err != nil {
			return nil, err
		}

		toType, err := getNodeType(conn.To.Type)
		if err != nil {
			return nil, err
		}

		connections = append(connections, domain.Connection{
			From: &domain.Node{
				ID:   conn.From.ID,
				Type: fromType,
			},
			To: &domain.Node{
				ID:   conn.To.ID,
				Type: toType,
			},
			Type: connType,
		})
	}
	return connections, nil
}

func getNodeType(nodeType string) (domain.NodeType, error) {
	if nodeType == "BLOCK" {
		return domain.BLOCK, nil
	} else if nodeType == "WORKFLOW" {
		return domain.WORKFLOW, nil
	} else {
		return domain.UNKNOWN_NODE_TYPE, fmt.Errorf("Unsupported connection type: %s", nodeType)
	}
}

func getConnectionType(connType string) (domain.ConnectionType, error) {
	if connType == "TRIGGERS" {
		return domain.TRIGGERS, nil
	} else if connType == "STARTS_WITH" {
		return domain.STARTS_WITH, nil
	} else if connType == "ENDS_WITH" {
		return domain.ENDS_WITH, nil
	} else {
		return domain.UNKNOWN_CONNECTION_TYPE, fmt.Errorf("Unsupported connection type: %s", connType)
	}
}

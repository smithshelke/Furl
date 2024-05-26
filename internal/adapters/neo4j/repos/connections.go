package repos

import (
	"fmt"

	"github.com/smithshelke/flur/internal/core/domain"
)

func getConnectionType(connType domain.ConnectionType) (string, error) {
	if connType == domain.TRIGGERS {
		return "TRIGGERS", nil
	} else if connType == domain.STARTS_WITH {
		return "STARTS_WITH", nil
	} else if connType == domain.ENDS_WITH {
		return "ENDS_WITH", nil
	} else {
		return "", fmt.Errorf("Unsupported connection type: %d", connType)
	}
}

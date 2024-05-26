package connection

import (
	"encoding/json"
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/web/common"
	"github.com/smithshelke/flur/internal/core/graphs"
)

type ConnectionController struct {
	graphAPI *graphs.GraphsAPI
}

func NewConnectionController(graphAPI *graphs.GraphsAPI) *ConnectionController {
	return &ConnectionController{
		graphAPI: graphAPI,
	}
}

func (c *ConnectionController) Create(w http.ResponseWriter, r *http.Request) {
	request := &CreateConnectionRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}

	connections, err := CreateConnectionRequestToConnections(request)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}

	err = c.graphAPI.CreateConnections(r.Context(), connections)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}

	response := &CreateConnectionRespone{}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

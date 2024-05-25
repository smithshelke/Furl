package web

import (
	"encoding/json"
	"net/http"
	"time"
)

func health(w http.ResponseWriter, req *http.Request) {
	response, _ := json.Marshal(&SuccessResponse{
		Message: "passed",
		Data:    time.Now(),
	})
	w.Write(response)
}

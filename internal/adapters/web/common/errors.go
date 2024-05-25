package common

import (
	"encoding/json"
	"net/http"
)

const (
	BAD_REQUEST          = 400
	SOMETHING_WENT_WRONG = 500
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func WriteError(w http.ResponseWriter, err error, code int) {
	errorResponse := &ErrorResponse{
		Message: err.Error(),
		Code:    400,
	}
	bytes, _ := json.Marshal(errorResponse)
	w.Write(bytes)
}

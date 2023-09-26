package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func Status(w http.ResponseWriter, status int, errorResponse interface{}) {
	response := Response{
		Status: StatusOK,
	}
	if errorResponse != nil {
		response.Status = StatusError
		response.Error = errorResponse
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(responseJson)
	if err != nil {
		return
	}
}

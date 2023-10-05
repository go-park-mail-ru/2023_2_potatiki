package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	StatusError = "Error"
)

type Response struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error,omitempty"`
}

func Err(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func JSON(w http.ResponseWriter, status int, response any) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseJson)))
	w.WriteHeader(status)
	w.Write(responseJson)
}

func JSONStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", "2")
	w.WriteHeader(status)
	w.Write([]byte("{}"))
}

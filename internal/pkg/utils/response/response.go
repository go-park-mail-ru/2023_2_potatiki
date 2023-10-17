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
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseJSON)))
	w.WriteHeader(status)
	_, err = w.Write(responseJSON)
	if err != nil {
		return // TODO: handle error
	}
}

func JSONStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json") // del
	w.Header().Set("Content-Length", "2")              // del
	w.WriteHeader(status)
	if _, err := w.Write([]byte("{}")); err != nil { // del
		return // TODO: handle error
	}
}

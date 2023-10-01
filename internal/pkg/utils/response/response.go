package response

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error,omitempty"`
}

const (
	//StatusOK    = "OK"
	StatusError = "Error"
)

//func OK() Response {
//	return Response{
//		Status: StatusOK,
//	}
//}

func Err(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func JSON(w http.ResponseWriter, status int, response any) {
	if response == nil {
		w.WriteHeader(status)
		log.Println("response.Status: response is nil")
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(status)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Засунуть длину в хеддер статус/err
	w.Header().Set("Content-Length", strconv.Itoa(len(responseJson)))
	w.WriteHeader(status)
	w.Write(responseJson) //err:uncheck
}

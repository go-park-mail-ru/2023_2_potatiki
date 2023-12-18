package responser

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
)

const (
	StatusError = "Error"
)

type response struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error,omitempty"`
}

func Err(msg string) *response {
	return &response{
		Status: StatusError,
		Error:  msg,
	}
}

type responser interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

func JSON(w http.ResponseWriter, status int, response responser) {
	responseJSON, err := response.MarshalJSON()
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
	w.WriteHeader(status)
}

func BodyErr(err error, log *slog.Logger, w http.ResponseWriter) bool {
	if err != nil {
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			JSON(w, http.StatusBadRequest, Err("request body is empty"))

			return true
		}
		log.Error("failed to decode request body", sl.Err(err))
		JSON(w, http.StatusBadRequest, Err("invalid request body"))

		return true
	}
	log.Debug("request body decoded")

	return false
}

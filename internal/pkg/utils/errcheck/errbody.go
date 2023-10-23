package errcheck

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"io"
	"log/slog"
	"net/http"
)

func BodyErr(err error, log *slog.Logger, w http.ResponseWriter) bool {
	if err != nil {
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))

			return true
		}
		log.Error("failed to decode request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusBadRequest)

		return true
	}

	return false
}

package csrfmw

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"

	"github.com/gorilla/mux"
)

const HEADER_NAME = "X-CSRF-Token"

func New(log *slog.Logger, jwtCORS jwter.JWTer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler { // TODO: del
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {
			case http.MethodGet:

				token, _, err := jwtCORS.EncodeCSRFToken(r.UserAgent())
				if err != nil {
					log.Error("error happened in Auther.GenerateToken", sl.Err(err))
					resp.JSONStatus(w, http.StatusUnauthorized)

					return
				}
				w.Header().Set(HEADER_NAME, token)

				return
			case http.MethodPost:
				token := r.Header.Get(HEADER_NAME)

				if token == "" {
					log.Error("miss csrf jwt")
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}
				//log.Debug("CSRF MW get csrf token", "token", token)

				UserAgent, err := jwtCORS.DecodeCSRFToken(token)

				if err != nil {
					log.Error("jws token is invalid csrf", sl.Err(err))
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}
				if r.UserAgent() != UserAgent {
					log.Error("UserAgent from token does not match request UserAgent", "UserAgent", UserAgent)
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}
				next.ServeHTTP(w, r)
			}
		})
	}
}

package csrfmw

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

const HEADER_NAME = "X-CSRF-Token"

func New(log *slog.Logger, jwtCORS jwter.JWTer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler { // TODO: del
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Id, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
			if !ok {
				log.Error("failed cast uuid from context value")
				resp.JSONStatus(w, http.StatusUnauthorized)

				return
			}

			switch r.Method {
			case http.MethodGet:

				token, _, err := jwtCORS.GenerateToken(&models.Profile{Id: Id})
				if err != nil {
					log.Error("error happened in Auther.GenerateToken", sl.Err(err))
					resp.JSONStatus(w, http.StatusUnauthorized)

					return
				}

				r.Header.Set(HEADER_NAME, token)

				return
			case http.MethodPost:
				token := r.Header.Get(HEADER_NAME)

				log.Debug("get csrf toke", "token", token)

				if token == "" {
					log.Error("miss csrf jwt")
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}

				claims, err := jwtCORS.GetClaims(token)
				if err != nil {
					log.Error("jws token is invalid csrf", sl.Err(err))
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}

				if Id != claims.ID {
					log.Error("jwt auth id does not match jwt csrf id", sl.Err(err))
					resp.JSONStatus(w, http.StatusForbidden)

					return
				}

				next.ServeHTTP(w, r)
			}
		})
	}
}

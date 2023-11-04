package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"

	"github.com/gorilla/mux"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Recover(log *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("Handle panic, recovered",
						slog.String("recover error", fmt.Sprintf("%v", err)),
						slog.String("url", r.URL.Path))
					resp.JSONStatus(w, http.StatusTooManyRequests)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

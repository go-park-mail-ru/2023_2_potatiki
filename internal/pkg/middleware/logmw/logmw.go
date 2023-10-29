package logmw

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const RequestIDCtx = "x-request-id"

type Config struct {
	DefaultLevel     slog.Level
	ClientErrorLevel slog.Level
	ServerErrorLevel slog.Level

	WithRequestID bool
}

// New returns a emux.MiddlewareFunc (middleware) that logs requests using slog.
//
// Requests with errors are logged using slog.Error().
// Requests without errors are logged using slog.Info().
func New(logger *slog.Logger) mux.MiddlewareFunc {
	return NewWithConfig(logger, Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithRequestID: true,
	})
}

type ResponseWrapper struct {
	http.ResponseWriter
	Status   int
	bytesLen int
}

func (r *ResponseWrapper) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *ResponseWrapper) Write(bytes []byte) (int, error) {
	r.bytesLen = len(bytes)

	return r.ResponseWriter.Write(bytes) //nolint:wrapcheck
}

func NewWithConfig(log *slog.Logger, config Config) mux.MiddlewareFunc { //nolint:cyclop
	return func(next http.Handler) http.Handler { // TODO: del
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapper := &ResponseWrapper{
				ResponseWriter: w,
				Status:         200,
			}

			requestID := r.Header.Get(RequestIDCtx) // TODO wrap
			if requestID == "" {
				requestID = uuid.New().String()
			}
			wrapper.Header().Set(RequestIDCtx, requestID)

			// ctx := context.WithValue(r.Context(), "request-id", requestID)
			r.Header.Set(RequestIDCtx, requestID)
			next.ServeHTTP(wrapper, r)

			status := wrapper.Status
			byteLen := wrapper.bytesLen

			// w.Header().Set(requestIDCtx, requestID)

			attributes := []slog.Attr{
				slog.Time("time", time.Now()),
				slog.String("duration", time.Since(start).String()),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", status),
				slog.String("remote-ip", r.RemoteAddr),
				slog.Int("bytes", byteLen),
				slog.String("user-agent", r.UserAgent()),
			}
			if config.WithRequestID {
				attributes = append(attributes, slog.String("request-id", requestID))
			}

			switch {
			case status >= http.StatusInternalServerError:
				log.LogAttrs(r.Context(), config.ServerErrorLevel, "Server Error", attributes...)
			case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
				log.LogAttrs(r.Context(), config.ClientErrorLevel, "Client Error", attributes...)
			case status >= http.StatusMultipleChoices && status < http.StatusBadRequest:
				log.LogAttrs(r.Context(), config.DefaultLevel, "Redirection", attributes...)
			case status >= http.StatusOK && status < http.StatusMultipleChoices:
				log.LogAttrs(r.Context(), config.DefaultLevel, "Success", attributes...)
			default:
				log.LogAttrs(r.Context(), config.DefaultLevel, "Informational", attributes...)
			}
		})
	}
}

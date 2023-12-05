package logmw

import (
	"bufio"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/metrics"
	"log/slog"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

const RequestIDCtx = "x-request-id"

var (
	UUIDRegExp = regexp.MustCompile(`[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`)
)

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
func New(mt *metrics.MetricHTTP, logger *slog.Logger) mux.MiddlewareFunc {
	return NewWithConfig(mt, logger, Config{
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

func (w *ResponseWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}

func (r *ResponseWrapper) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *ResponseWrapper) Write(bytes []byte) (int, error) {
	r.bytesLen = len(bytes)

	return r.ResponseWriter.Write(bytes) //nolint:wrapcheck
}

func NewWithConfig(mt *metrics.MetricHTTP, log *slog.Logger, config Config) mux.MiddlewareFunc { //nolint:cyclop
	return func(next http.Handler) http.Handler { // TODO: del
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			wrapper := &ResponseWrapper{
				ResponseWriter: w,
				Status:         200,
			}

			requestID := r.Header.Get(RequestIDCtx) // TODO wrap
			if requestID == "" {
				requestID = uuid.NewV4().String()
			}
			wrapper.Header().Set(RequestIDCtx, requestID)

			// ctx := context.WithValue(r.Context(), "request-id", requestID)
			r.Header.Set(RequestIDCtx, requestID)

			start := time.Now()
			next.ServeHTTP(wrapper, r)
			duration := time.Since(start)

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
			// r.Response.Status
			bytesUrl := []byte(r.URL.Path)
			urlWithCuttedUUID := UUIDRegExp.ReplaceAll(bytesUrl, []byte("<uuid>"))

			switch {
			case status >= http.StatusInternalServerError:
				mt.IncreaseErr(strconv.Itoa(status), string(urlWithCuttedUUID), r.Method)
				log.LogAttrs(r.Context(), config.ServerErrorLevel, "Server Error", attributes...)
			case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
				mt.IncreaseErr(strconv.Itoa(status), string(urlWithCuttedUUID), r.Method)
				log.LogAttrs(r.Context(), config.ClientErrorLevel, "Client Error", attributes...)
			case status >= http.StatusMultipleChoices && status < http.StatusBadRequest:
				mt.IncreaseErr(strconv.Itoa(status), string(urlWithCuttedUUID), r.Method)
				log.LogAttrs(r.Context(), config.DefaultLevel, "Redirection", attributes...)
			case status >= http.StatusOK && status < http.StatusMultipleChoices:
				log.LogAttrs(r.Context(), config.DefaultLevel, "Success", attributes...)
			default:
				log.LogAttrs(r.Context(), config.DefaultLevel, "Informational", attributes...)
			}
			mt.IncreaseHits(string(urlWithCuttedUUID), r.Method)
			mt.AddDurationToHistogram(string(urlWithCuttedUUID), r.Method, duration)
			mt.AddDurationToSummary(strconv.Itoa(status), string(urlWithCuttedUUID), r.Method, duration)
		})
	}
}

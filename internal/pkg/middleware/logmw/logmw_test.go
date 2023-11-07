package logmw

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
)

func TestNew(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout))
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

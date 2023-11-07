package csrfmw

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
)

func TestNew_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().DecodeCSRFToken(gomock.Any()).Return("fb11fe90-09bb-4e72-98a5-5ffba93aa39a", nil)

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodPost, "http://www.your-domain.com", nil)

	req.Header.Set(HEADER_NAME, "2oo3hri3irj")

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_Post_Bad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().DecodeCSRFToken(gomock.Any()).Return("fb11fe90-09bb-4e72-98a5-5ffba93aa39a", errors.New("dummyError"))

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodPost, "http://www.your-domain.com", nil)

	req.Header.Set(HEADER_NAME, "2oo3hri3irj")

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_Post_BadHeader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockJWTer(ctrl)

	//uc.EXPECT().DecodeCSRFToken(gomock.Any()).Return("fb11fe90-09bb-4e72-98a5-5ffba93aa39a", errors.New("dummyError"))

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodPost, "http://www.your-domain.com", nil)

	req.Header.Set(HEADER_NAME, "")

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().EncodeCSRFToken(gomock.Any()).Return("fb11fe90-09bb-4e72-98a5-5ffba93aa39a", time.Now(), nil)

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	//req.Header.Set(HEADER_NAME, "2oo3hri3irj")

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_Get_Bad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().EncodeCSRFToken(gomock.Any()).Return("fb11fe90-09bb-4e72-98a5-5ffba93aa39a", time.Now(), errors.New("dummyError"))

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	//req.Header.Set(HEADER_NAME, "2oo3hri3irj")

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

package authmw

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
	uuid "github.com/satori/go.uuid"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	/*cfg := mock.NewMockConfiger(ctrl)
	cfg.EXPECT().GetIssuer().Return("auth")
	cfg.EXPECT().GetSecret().Return("aofuhorugugalohp30q94gpwg")
	cfg.EXPECT().GetTTL().Return(time.Hour * 6)*/

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().DecodeAuthToken(gomock.Any()).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	req.AddCookie(MakeTokenCookie("121231231233", time.Now().Add(time.Hour*6)))

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_BadCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	/*cfg := mock.NewMockConfiger(ctrl)
	cfg.EXPECT().GetIssuer().Return("auth")
	cfg.EXPECT().GetSecret().Return("aofuhorugugalohp30q94gpwg")
	cfg.EXPECT().GetTTL().Return(time.Hour * 6)*/

	uc := mock.NewMockJWTer(ctrl)

	//uc.EXPECT().DecodeAuthToken(gomock.Any()).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	//req.AddCookie(MakeTokenCookie("121231231233", time.Now().Add(time.Hour*6)))

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

func TestNew_BadDecode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	/*cfg := mock.NewMockConfiger(ctrl)
	cfg.EXPECT().GetIssuer().Return("auth")
	cfg.EXPECT().GetSecret().Return("aofuhorugugalohp30q94gpwg")
	cfg.EXPECT().GetTTL().Return(time.Hour * 6)*/

	uc := mock.NewMockJWTer(ctrl)

	uc.EXPECT().DecodeAuthToken(gomock.Any()).Return(uuid.UUID{}, errors.New("dummyError"))

	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	req.AddCookie(MakeTokenCookie("121231231233", time.Now().Add(time.Hour*6)))

	res := httptest.NewRecorder()
	handler(res, req)

	mw := New(logger.Set("local", os.Stdout), uc)
	mw(http.HandlerFunc(handler)).ServeHTTP(res, req)
}

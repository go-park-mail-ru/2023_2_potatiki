package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)
	id := uuid.NewV4()
	uc.EXPECT().GetProduct(gomock.Any(), id).Return(
		models.Product{
			Id:          id,
			Name:        "123",
			Description: "123",
			Price:       123,
		}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo",
		strings.NewReader(
			"{ \"id\": \""+id.String()+"\", \"name\": \"123\" , \"description\": \"123\", \"price\": \"123\"}"))
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})
	w := httptest.NewRecorder()
	ProductsHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
	ProductsHandler.Product(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)

	t.Run("EmptyID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		ProductHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
		ProductHandler.Product(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalidID"})
		w := httptest.NewRecorder()
		ProductHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
		ProductHandler.Product(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("GetProductError", func(t *testing.T) {
		validID := uuid.NewV4()
		uc.EXPECT().GetProduct(gomock.Any(), validID).Return(models.Product{}, errors.New("getProductError"))

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": validID.String()})
		w := httptest.NewRecorder()
		ProductHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
		ProductHandler.Product(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)
	id := uuid.NewV4()
	uc.EXPECT().GetProducts(gomock.Any(), int64(0), int64(1)).Return(
		[]models.Product{{
			Id:          id,
			Name:        "123",
			Description: "123",
			Price:       123,
		}}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo",
		strings.NewReader(
			"[{ \"id\": \""+id.String()+"\", \"name\": \"123\" , \"description\": \"123\", \"price\": \"123\"}]"))
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	ProductsHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
	ProductsHandler.Products(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductsBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)
	id := uuid.NewV4()

	uc.EXPECT().GetProducts(gomock.Any(), int64(0), int64(1)).Return(nil, errors.New("some error"))

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo",
		strings.NewReader(
			"[{ \"id\": \""+id.String()+"\", \"name\": \"123\" , \"description\": \"123\", \"price\": \"123\"}]"))
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	ProductsHandler := NewProductsHandler(logger.Set("prod", os.Stdout), uc)
	ProductsHandler.Products(w, req)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

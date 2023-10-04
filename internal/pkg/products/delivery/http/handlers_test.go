package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)
	id := uuid.New()
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
	ProductsHandler := NewProductsHandler(logger.Set("prod"), uc)
	ProductsHandler.Product(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsUsecase(ctrl)
	id := uuid.New()
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
	ProductsHandler := NewProductsHandler(logger.Set("prod"), uc)
	ProductsHandler.Products(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

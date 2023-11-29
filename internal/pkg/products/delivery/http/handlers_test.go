package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func TestProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.NewV4()

	client := mock.NewMockProductsClient(ctrl)

	client.EXPECT().GetProduct(gomock.Any(), &gen.ProductRequest{Id: id.String()}).Return(
		&gen.ProductResponse{
			Product: &gmodels.Product{
				Id:          id.String(),
				Name:        "123",
				Description: "123",
				Price:       123,
				Category:    &gmodels.Category{},
			}}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.String()})
	w := httptest.NewRecorder()
	ProductsHandler := NewProductsHandler(client, logger.Set("prod", os.Stdout))
	ProductsHandler.Product(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

/*
func TestProductBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsClient(ctrl)

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
*/
/*
func TestProductsHandledddr_Category(t *testing.T) {
	testCases := []struct {
		name               string
		MockProductsClient func(*mock.MockProductsClient)
		expectedStatus     int
		funcCtxUser        func(context.Context) context.Context
		param1             string
		param2             string
		param3             string
	}{
		{
			name: "SuccessfulCategory",
			MockProductsClient: func(mock *mock.MockProductsClient) {
				mock.EXPECT().GetCategory(gomock.Any(), &gen.CategoryRequest{}).
					Return(&gen.CategoryResponse{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "1",
			param2: "2",
			param3: "3",
		},
		{
			name: "UnsuccessfulCategoryWithCorrectQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {
				mock.EXPECT().GetCategory(gomock.Any(), &gen.CategoryRequest{}).
					Return(&gen.CategoryResponse{}, errors.New("error in get product by category"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "1",
			param2: "2",
			param3: "3",
		},
		{
			name:               "UnsuccessfulCategoryWithNotIntInPagingQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {},
			expectedStatus:     http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "1",
			param2: "qwerty",
			param3: "",
		},
		{
			name:               "UnsuccessfulCategoryWithNotIntInCountQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {},
			expectedStatus:     http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "1",
			param2: "2",
			param3: "qwerty",
		},
		{
			name:               "UnsuccessfulCategoryWithFirstQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {},
			expectedStatus:     http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "",
			param2: "",
			param3: "",
		},
		{
			name:               "UnsuccessfulCategoryWithSecondQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {},
			expectedStatus:     http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "",
			param2: "2",
			param3: "",
		},
		{
			name:               "UnsuccessfulCategoryWithThirdQuery",
			MockProductsClient: func(mock *mock.MockProductsClient) {},
			expectedStatus:     http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			param1: "",
			param2: "",
			param3: "3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock.NewMockProductsClient(ctrl)
			tc.MockProductsClient(mock)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/address/get_current", nil)

			q := req.URL.Query()
			q.Add("category_id", tc.param1)
			q.Add("paging", tc.param2)
			q.Add("count", tc.param3)
			req.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewProductsHandler(mock, logger.Set("local", os.Stdout))
			addressHandler.Category(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}


func TestProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProductsClient(ctrl)
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

	uc := mock.NewMockProductsClient(ctrl)

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

	uc := mock.NewMockProductsClient(ctrl)
	id := uuid.NewV4()
	uc.EXPECT().GetProducts(gomock.Any(), int64(0), int64(1), gomock.Any(), gomock.Any()).Return(
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

	uc := mock.NewMockProductsClient(ctrl)
	id := uuid.NewV4()

	uc.EXPECT().GetProducts(gomock.Any(), int64(0), int64(1), gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo",
		strings.NewReader(
			"[{ \"id\": \""+id.String()+"\", \"name\": \"123\" , \"description\": \"123\", \"price\": \"123\"}]"))
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	ProductsHandler := NewProductsHandler(logger.Set("prod", os.Stdout))
	ProductsHandler.Products(w, req)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}
*/

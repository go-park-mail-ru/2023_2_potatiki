package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCategoryHandler_Categories(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCategoryUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulCategories",
			mockUsecaseFn: func(mockUsecase *mock.MockCategoryUsecase) {
				mockUsecase.EXPECT().Categories(gomock.Any()).Return(models.CategoryTree{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "UnsuccessfulGetCart",
			mockUsecaseFn: func(mockUsecase *mock.MockCategoryUsecase) {
				mockUsecase.EXPECT().Categories(gomock.Any()).Return(models.CategoryTree{}, errors.New("error in get categories"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCategoryUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/category/get_all", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCategoryHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.Categories(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

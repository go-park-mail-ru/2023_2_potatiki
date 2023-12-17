package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestRecommendationsHandler_AnonRecommendations(t *testing.T) {
	tests := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockRecommendationsUsecase)
		expectedStatus int
		queryString    string
	}{
		{
			name: "SuccessfulAnonRecommendations",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().AnonRecommendations(
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(
					models.ProductSlice{},
					nil,
				)
			},
			expectedStatus: http.StatusOK,
			queryString:    `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=6`,
		},
		{
			name: "UsecaseErrorRecommendations",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().AnonRecommendations(
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(
					models.ProductSlice{},
					errors.New("error get recommendations"),
				)
			},
			expectedStatus: http.StatusTooManyRequests,
			queryString:    `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=6`,
		},
		{
			name:           "EmptyIdAnonRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			queryString:    `?category_id=6`,
		},
		{
			name:           "EmptyCategoryIdAnonRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			queryString:    `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d`,
		},
		{
			name:           "BadIdAnonRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			queryString:    `?id=ffbd6bc7-476d-44d2-805ef714a70d&category_id=6`,
		},
		{
			name:           "BadCategoryIdAnonRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			queryString:    `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=d`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockRecommendationsUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(
				http.MethodGet,
				"http://zuzu-market.ru/api/recommendations/get_anon"+tc.queryString,
				nil,
			)
			w := httptest.NewRecorder()

			h := NewRecommendationsHandler(logger.Set("local", os.Stdout), mockUsecase)
			h.AnonRecommendations(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestRecommendationsHandler_Recommendations(t *testing.T) {
	tests := []struct {
		name           string
		mockUsecaseFn  func(usecase *mock.MockRecommendationsUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		queryString    string
	}{
		{
			name: "SuccessfulRecommendations",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().Recommendations(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(models.ProductSlice{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=6`,
		},
		{
			name: "SuccessfulRecommendations",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().Recommendations(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(models.ProductSlice{}, errors.New("error get recommendations"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=6`,
		},
		{
			name:           "EmptyIdRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?category_id=6`,
		},
		{
			name:           "BadIdRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-`,
		},
		{
			name:           "EmptyCategoryIdRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d`,
		},
		{
			name:           "BadCategoryIdRecommendations",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=8d`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			queryString: `?id=ffbd6bc7-476d-44d2-a08b-805ef714a70d&category_id=8`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockRecommendationsUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(
				http.MethodGet,
				"http://zuzu-market.ru/api/recommendations/get_anon"+tc.queryString,
				nil,
			)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)

			h := NewRecommendationsHandler(logger.Set("local", os.Stdout), mockUsecase)
			h.Recommendations(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestRecommendationsHandler_UpdateUserActivity(t *testing.T) {
	tests := []struct {
		name           string
		mockUsecaseFn  func(usecase *mock.MockRecommendationsUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulUpdateUserActivity",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().UpdateUserActivity(
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{"product": [
				{
					"productId": "ffbd6bc7-476d-44d2-a08b-805ef714a70d",
					"activityPoints": 1,
					"isBought": true,
					"isReviewed": true
				},
				{
					"productId": "ffb35880-f073-479d-88e5-4557407d3fd6",
					"activityPoints": 2
				}
			]
			}`,
		},
		{
			name: "UsecaseErrorUpdateUserActivity",
			mockUsecaseFn: func(mockUsecase *mock.MockRecommendationsUsecase) {
				mockUsecase.EXPECT().UpdateUserActivity(
					gomock.Any(),
					gomock.Any(),
					gomock.Any()).Return(errors.New("error in update activity"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{"product": [
				{
					"productId": "ffbd6bc7-476d-44d2-a08b-805ef714a70d",
					"activityPoints": 1,
					"isBought": true,
					"isReviewed": true
				},
				{
					"productId": "ffb35880-f073-479d-88e5-4557407d3fd6",
					"activityPoints": 2
				}
			]
			}`,
		},
		{
			name:           "UserUnauthorized",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{"product": [
				{
					"productId": "ffbd6bc7-476d-44d2-a08b-805ef714a70d",
					"activityPoints": 1,
					"isBought": true,
					"isReviewed": true
				},
				{
					"productId": "ffb35880-f073-479d-88e5-4557407d3fd6",
					"activityPoints": 2
				}
			]
			}`,
		},
		{
			name:           "BadBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockRecommendationsUsecase) {},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{"product": [
				{
					"productId": "ffbd6bc7-476d-44d2-a08b-805ef714a70d",
					"activityPoints": 1,
					"isBought": true,
					"isReviewed": true
				
				{
					"productId": "ffb35880-f073-479d-88e5-4557407d3fd6",
					"activityPoints": 2
				}
			]
			}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockRecommendationsUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(
				http.MethodGet,
				"http://zuzu-market.ru/api/recommendations/get_anon",
				strings.NewReader(tc.jsonString),
			)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)

			h := NewRecommendationsHandler(logger.Set("local", os.Stdout), mockUsecase)
			h.UpdateUserActivity(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

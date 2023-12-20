package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/comments/mocks"
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

func TestCommentsHandler_CreateComment(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCommentsUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulCreateComment",
			mockUsecaseFn: func(mockUsecase *mock.MockCommentsUsecase) {
				mockUsecase.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(models.Comment{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"userName": "JohnDoe",
				"creationDate": "2023-12-19T15:30:00Z",
				"productId": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"pros": "High quality, durable",
				"cons": "Expensive",
				"comment": "This product is great, but it's a bit pricey.",
				"rating": 4
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockCommentsUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockCommentsUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulCreateComment",
			mockUsecaseFn: func(mockUsecase *mock.MockCommentsUsecase) {
				mockUsecase.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(models.Comment{}, errors.New("error in create comment"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"userName": "JohnDoe",
				"creationDate": "2023-12-19T15:30:00Z",
				"productId": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"pros": "High quality, durable",
				"cons": "Expensive",
				"comment": "This product is great, but it's a bit pricey.",
				"rating": 4
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCommentsUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/comments/create",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCommentsHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.CreateComment(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestCommentsHandler_GetProductComments(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCommentsUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
		hasQuery       bool
	}{
		{
			name: "SuccessfulGetProductComments",
			mockUsecaseFn: func(mockUsecase *mock.MockCommentsUsecase) {
				mockUsecase.EXPECT().GetProductComments(gomock.Any(), gomock.Any()).Return([]models.Comment{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"userName": "JohnDoe",
				"creationDate": "2023-12-19T15:30:00Z",
				"productId": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"pros": "High quality, durable",
				"cons": "Expensive",
				"comment": "This product is great, but it's a bit pricey.",
				"rating": 4
			}`,
			hasQuery: true,
		},
		{
			name: "EmptyQuery",
			mockUsecaseFn: func(mockUsecase *mock.MockCommentsUsecase) {
			},
			expectedStatus: http.StatusBadRequest,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"userName": "JohnDoe",
				"creationDate": "2023-12-19T15:30:00Z",
				"productId": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"pros": "High quality, durable",
				"cons": "Expensive",
				"comment": "This product is great, but it's a bit pricey.",
				"rating": 4
			}`,
			hasQuery: false,
		},
		{
			name: "UnsuccessfulGetProductComments",
			mockUsecaseFn: func(mockUsecase *mock.MockCommentsUsecase) {
				mockUsecase.EXPECT().GetProductComments(gomock.Any(), gomock.Any()).Return([]models.Comment{}, errors.New("error in get comments"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"userName": "JohnDoe",
				"creationDate": "2023-12-19T15:30:00Z",
				"productId": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"pros": "High quality, durable",
				"cons": "Expensive",
				"comment": "This product is great, but it's a bit pricey.",
				"rating": 4
			}`,
			hasQuery: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCommentsUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/comments/get_all",
				strings.NewReader(tc.jsonString))

			if tc.hasQuery {
				q := req.URL.Query()
				q.Add("product", "6ba7b810-9dad-11d1-80b4-00c04fd430c8")
				req.URL.RawQuery = q.Encode()
			}

			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCommentsHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetProductComments(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

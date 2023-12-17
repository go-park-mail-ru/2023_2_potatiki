package repo

import (
	"context"
	"errors"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRecommendationsRepo_CreateUserActivity(t *testing.T) {
	userID := uuid.NewV4()
	ErrCreateActivity := errors.New("error create user activity")
	tests := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		columns    []string
		id         uuid.UUID
		err        error
	}{
		{
			name: "SuccessfulCreateUserActivity",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					createActivity,
					userID,
				).Return(nil, nil)
			},
			columns: []string{"user_id"},
			id:      userID,
			err:     nil,
		},
		{
			name: "FailureCreateUserActivity",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					createActivity,
					userID,
				).Return(nil, ErrCreateActivity)
			},
			columns: []string{"user_id"},
			id:      userID,
			err:     ErrCreateActivity,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewRecommendationsRepo(mockPool)
			err := repo.CreateUserActivity(context.Background(), tc.id)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRecommendationsRepo_ReadProduct(t *testing.T) {
	uuidForTest := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         uuid.UUID
		paging     int64
		count      int64
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadProduct",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), getProduct, uuidForTest).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      uuidForTest,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", int64(0), "", float32(0), int64(0), "", int64(0)).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewRecommendationsRepo(mockPool)
			_, err := repo.ReadProduct(context.Background(), tc.id)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestRecommendationsRepo_ReadProductsFromCategory(t *testing.T) {
	tests := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int64
		err        error
	}{
		{
			name: "SuccessfulReadProductsFromCategory",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getProductsFromCategory,
					int64(6),
					models.ProductCountFromCategory,
				).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      int64(6),
			err:     nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", int64(0), "", float32(0), int64(0), "", int64(0)).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewRecommendationsRepo(mockPool)
			_, err := repo.ReadProductsFromCategory(context.Background(), tc.id)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRecommendationsRepo_ReadProductsFromParentCategory(t *testing.T) {
	tests := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int64
		err        error
	}{
		{
			name: "SuccessfulReadProductsFromParentCategory",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getProductsFromParentCategory,
					int64(6),
				).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      int64(6),
			err:     nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", int64(0), "", float32(0), int64(0), "", int64(0)).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewRecommendationsRepo(mockPool)
			_, err := repo.ReadProductsFromParentCategory(context.Background(), tc.id)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRecommendationsRepo_ReadUserActivity(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.UserActivityStore
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecommendationsRepo{
				db: tt.fields.db,
			}
			got, err := r.ReadUserActivity(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUserActivity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadUserActivity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecommendationsRepo_UpdateUserActivity(t *testing.T) {
	userID := uuid.NewV4()
	ErrUpdateActivity := errors.New("error update user activity")
	tests := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		id         uuid.UUID
		err        error
	}{
		{
			name: "SuccessfullUpdateUserActivity",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUserActivity,
					userID,
					gomock.Any(),
				).Return(nil, nil)
			},
			id:  userID,
			err: nil,
		},
		{
			name: "FailureUpdateUserActivity",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUserActivity,
					userID,
					gomock.Any(),
				).Return(nil, ErrUpdateActivity)
			},
			id:  userID,
			err: ErrUpdateActivity,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewRecommendationsRepo(mockPool)
			err := repo.UpdateUserActivity(context.Background(), tc.id, models.UserActivityStore{})

			assert.ErrorIs(t, err, tc.err)
		})
	}
}

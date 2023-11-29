package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestOrderRepo_CreateOrder(t *testing.T) {
	addressID := uuid.NewV4()
	userID := uuid.NewV4()
	var statusID int64 = 0
	tests := []struct {
		name       string
		mockRepoFn func(pool *pgxpoolmock.MockPgxPool)
		err        error
	}{{
		name: "SuccessfullCreateOrder",
		mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
			mockPool.EXPECT().Exec(
				gomock.Any(),
				createOrder,
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any()).Return(nil, nil)
		},
		err: nil,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tt.mockRepoFn(mockPool)

			repo := NewOrderRepo(mockPool)
			_, err := repo.CreateOrder(context.Background(), models.Cart{}, addressID, userID, statusID)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestOrderRepo_ReadOrderID(t *testing.T) {
	userID := uuid.NewV4()
	tests := []struct {
		name       string
		mockRepoFn func(pool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows)
		err        error
		columns    []string
	}{
		{
			name: "SuccessfullReadOrderID",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(
					gomock.Any(),
					getCurrentOrderID,
					gomock.Any(),
				).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"order_id"},
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tt.columns).AddRow(uuid.UUID{}).ToPgxRows()
			tt.mockRepoFn(mockPool, pgxRows)

			repo := NewOrderRepo(mockPool)
			_, err := repo.ReadOrderID(context.Background(), userID)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestOrderRepo_ReadOrder(t *testing.T) {
	userID := uuid.NewV4()
	tests := []struct {
		name       string
		mockRepoFn func(pool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows)
		err        error
		columns    []string
	}{
		{
			name: "SuccessfullReadOrder",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getCurrentOrder,
					gomock.Any(),
				).Return(pgxRows, nil)
			},
			columns: []string{
				"product_id",
				"product_name",
				"product_description",
				"product_price",
				"product_imgsrc",
				"product_rating",
				"product_quantity",
				"category_id",
				"category_name",
				"status_id",
				"address_id",
				"address_city",
				"address_street",
				"address_house",
				"address_flat",
				"is_current",
			},
			err: nil,
		},
		{
			name: "UnSuccessfullReadOrder",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getCurrentOrder,
					gomock.Any(),
				).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{
				"product_id",
				"product_name",
				"product_description",
				"product_price",
				"product_imgsrc",
				"product_rating",
				"product_quantity",
				"category_id",
				"category_name",
				"status_id",
				"address_id",
				"address_city",
				"address_street",
				"address_house",
				"address_flat",
				"is_current",
			},
			err: ErrPoductsInOrderNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tt.columns).AddRow(
				uuid.UUID{},
				"",
				"",
				int64(0),
				"",
				float32(0),
				int64(0),
				int64(0),
				"",
				int64(0),
				uuid.UUID{},
				"",
				"",
				"",
				"",
				true,
			).ToPgxRows()
			tt.mockRepoFn(mockPool, pgxRows)

			repo := NewOrderRepo(mockPool)
			_, err := repo.ReadOrder(context.Background(), userID)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestOrderRepo_ReadOrdersID(t *testing.T) {
	userID := uuid.NewV4()
	tests := []struct {
		name       string
		mockRepoFn func(pool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows)
		err        error
		columns    []string
	}{
		{
			name: "SuccessfullReadOrdersID",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getOrdersID,
					gomock.Any(),
				).Return(pgxRows, nil)
			},
			columns: []string{
				"order_id",
			},
			err: nil,
		},
		{
			name: "UnSuccessfullReadOrdersID",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(
					gomock.Any(),
					getOrdersID,
					gomock.Any(),
				).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{
				"order_id",
			},
			err: ErrOrdersNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tt.columns).AddRow(uuid.UUID{}).ToPgxRows()
			tt.mockRepoFn(mockPool, pgxRows)

			repo := NewOrderRepo(mockPool)
			_, err := repo.ReadOrdersID(context.Background(), userID)

			assert.Equal(t, tt.err, err)
		})
	}
}

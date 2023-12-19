package repo

import (
	"context"
	"fmt"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestCartRepo_ReadCartProducts(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadCartProducts",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgSrc", "rating", "quantity", "Category.ID", "Category.Name"},
			err:     nil,
		},
		{
			name: "UnsuccessfulReadCartProductsErrProductsNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "description", "price", "imgSrc", "rating", "quantity", "Category.ID", "Category.Name"},
			err:     ErrProductsNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", int64(0), "", float32(0), int64(0), int64(0), "").ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewCartRepo(mockPool)
			_, err := repo.ReadCartProducts(context.Background(), models.Cart{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCartRepo_CreateCart(t *testing.T) {
	cartID := uuid.NewV4()
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
		id         uuid.UUID
	}{
		{
			name: "SuccessfulCreateCart",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), createCart, gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			err: nil,
			id:  cartID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewCartRepo(mockPool)
			_, err := repo.CreateCart(context.Background(), userID)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.id, cartID)
		})
	}
}

func TestCartRepo_DeleteCart(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulDeleteCart",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), deleteCard, gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewCartRepo(mockPool)
			err := repo.DeleteCart(context.Background(), userID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCartRepo_CheckCart(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulCheckCart",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), getCart, userID).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"id"},
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewCartRepo(mockPool)
			_, err := repo.CheckCart(context.Background(), userID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCartRepo_AddProduct(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "UnsuccessfulAddProduct",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), updateOrCreateProduct, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, pgx.ErrNoRows)
			},
			err: fmt.Errorf("error happened in db.Exec: %w", pgx.ErrNoRows),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewCartRepo(mockPool)
			_, err := repo.AddProduct(context.Background(), models.Cart{}, models.CartProductUpdate{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCartRepo_DeleteProduct(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "UnsuccessfulDeleteProduct",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), deleteProduct, gomock.Any(), gomock.Any()).Return(nil, pgx.ErrNoRows)
			},
			err: fmt.Errorf("error happened in db.Exec: %w", pgx.ErrNoRows),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewCartRepo(mockPool)
			_, err := repo.DeleteProduct(context.Background(), models.Cart{}, models.CartProductDelete{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCartRepo_ReadCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.ReadCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.ReadCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_UpdateCart(t *testing.T) {
	type args struct {
		ctx  context.Context
		cart models.Cart
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UpdateCart(tt.args.ctx, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.UpdateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.UpdateCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

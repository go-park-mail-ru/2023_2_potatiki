package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// mock on repo and uid used as uuid for everything that has any uid filed
// func setUp(t *testing.T, pgxRows pgx.Rows) *ProductsRepo {
//	os.Setenv("SECRET", "test")
//	ctl := gomock.NewController(t)
//	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
//
//	repo := NewAuthRepo(mockPool, nil)
//	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
//	return repo
//}

func TestProductsRepo_ReadProducts(t *testing.T) {
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating", "Category.Id", "Category.Name"}
	id := uuid.UUID{}
	repo := NewProductsRepo(mockPool)
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(id, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, nil)

	products, err := repo.ReadProducts(context.Background(), 0, 10)

	assert.Nil(t, err)
	assert.Len(t, products, 1)
}

func TestProductsRepo_ReadProduct(t *testing.T) {
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating", "Category.Id", "Category.Name"}
	id := uuid.NewV4()
	repo := NewProductsRepo(mockPool)
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(id, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
	mockPool.EXPECT().QueryRow(gomock.Any(), getProduct, gomock.Any()).Return(pgxRows)
	pgxRows.Next()

	_, err := repo.ReadProduct(context.Background(), id)

	assert.Nil(t, err)
	// assert.Len(t, products, 1)
}

func TestProductsRepo_ReadCategory(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		err        error
	}{
		{
			name: "SuccessfulReadCategory",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByCategoryID, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name"},
			id:      2,
			paging:  4,
			count:   6,
			err:     nil,
		},
		{
			name: "UnsuccessfulReadCategoryProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByCategoryID, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name"},
			id:      2,
			paging:  4,
			count:   6,
			err:     ErrPoductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadCategory(context.Background(), tc.id, tc.paging, tc.count)

			assert.Equal(t, tc.err, err)
		})
	}
	//ctr := gomock.NewController(t)
	//mockPool := pgxpoolmock.NewMockPgxPool(ctr)
	//defer ctr.Finish()
	//
	//columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating", "Category.Id", "Category.Name"}
	//pgxRows := pgxpoolmock.NewRows(columns).
	//	AddRow(id, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
	//mockPool.EXPECT().QueryRow(gomock.Any(), getProduct, gomock.Any()).Return(pgxRows)
	//pgxRows.Next()
	//
	//repo := NewProductsRepo(mockPool)
	//_, err := repo.ReadProduct(context.Background(), id)
	//
	//assert.Nil(t, err)
	// assert.Len(t, products, 1)
}

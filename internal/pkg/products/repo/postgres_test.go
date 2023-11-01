package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
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

	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating"}
	id := uuid.UUID{}
	repo := NewProductsRepo(mockPool)
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(id, "", "", int64(0), "", float64(0)).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, nil)

	products, err := repo.ReadProducts(context.Background(), 0, 10)

	assert.Nil(t, err)
	assert.Len(t, products, 1)
}

func TestProductsRepo_ReadProduct(t *testing.T) {
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating"}
	id := uuid.NewV4()
	repo := NewProductsRepo(mockPool)
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(id, "", "", int64(0), "", float64(0)).ToPgxRows()
	mockPool.EXPECT().QueryRow(gomock.Any(), getProduct, gomock.Any()).Return(pgxRows)
	pgxRows.Next()

	_, err := repo.ReadProduct(context.Background(), id)

	assert.Nil(t, err)
	// assert.Len(t, products, 1)
}

package repo

//
//import (
//	"context"
//	"github.com/driftprogramming/pgxpoolmock"
//	"github.com/golang/mock/gomock"
//	"github.com/jackc/pgx/v4"
//	"github.com/satori/go.uuid"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestProductsRepo_ReadCategory(t *testing.T) {
//	testCases := []struct {
//		name       string
//		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
//		columns    []string
//		id         int
//		paging     int64
//		count      int64
//		err        error
//	}{
//		{
//			name: "SuccessfulReadCategory",
//			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
//				mockPool.EXPECT().Query(gomock.Any(), getProductsByCategoryID, gomock.Any()).Return(pgxRows, nil)
//			},
//			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name"},
//			id:      2,
//			paging:  4,
//			count:   6,
//			err:     nil,
//		},
//		{
//			name: "UnsuccessfulReadCategoryProductNotFound",
//			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
//				mockPool.EXPECT().Query(gomock.Any(), getProductsByCategoryID, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
//			},
//			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name"},
//			id:      2,
//			paging:  4,
//			count:   6,
//			err:     ErrProductNotFound,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			ctr := gomock.NewController(t)
//			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
//			defer ctr.Finish()
//
//			pgxRows := pgxpoolmock.NewRows(tc.columns).
//				AddRow(uuid.UUID{}, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
//
//			tc.mockRepoFn(mockPool, pgxRows)
//
//			repo := NewProductsRepo(mockPool)
//			_, err := repo.ReadCategory(context.Background(), tc.id, tc.paging, tc.count, "DESC", "DESC")
//
//			assert.Equal(t, tc.err, err)
//		})
//	}
//}
//
//func TestProductsRepo_ReadProducts(t *testing.T) {
//	ctl := gomock.NewController(t)
//	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
//
//	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating", "Category.Id", "Category.Name"}
//	id := uuid.UUID{}
//	repo := NewProductsRepo(mockPool)
//	pgxRows := pgxpoolmock.NewRows(columns).
//		AddRow(id, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
//	mockPool.EXPECT().Query(gomock.Any(), getProductsDESCRatingPrice, gomock.Any()).Return(pgxRows, nil)
//
//	products, err := repo.ReadProducts(context.Background(), 0, 10, "DESC", "DESC")
//
//	assert.Nil(t, err)
//	assert.Len(t, products, 1)
//}
//
//func TestProductsRepo_ReadProduct(t *testing.T) {
//	ctl := gomock.NewController(t)
//	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
//
//	columns := []string{"Id", "NameProduct", "Description", "Price", "ImgSrc", "Rating", "Category.Id", "Category.Name"}
//	id := uuid.NewV4()
//	repo := NewProductsRepo(mockPool)
//	pgxRows := pgxpoolmock.NewRows(columns).
//		AddRow(id, "", "", int64(0), "", float64(0), int64(0), "").ToPgxRows()
//	mockPool.EXPECT().QueryRow(gomock.Any(), getProduct, gomock.Any()).Return(pgxRows)
//	pgxRows.Next()
//
//	_, err := repo.ReadProduct(context.Background(), id)
//
//	assert.Nil(t, err)
//}

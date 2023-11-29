package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductsRepo_ReadProductsCategory(t *testing.T) {
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
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
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
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			err:     ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsCategory(context.Background(), tc.id, tc.paging, tc.count)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsCategoryByPrice(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadProductsByPrice",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByPrice, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     nil,
		},
		{
			name: "UnsuccessfulReadProductsByPriceProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByPrice, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsCategoryByPrice(context.Background(), tc.id, tc.paging, tc.count, tc.priceBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsCategoryByRating(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		ratingBy   string
		err        error
	}{
		{
			name: "SuccessfulReadProductsByPrice",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByRating, gomock.Any()).Return(pgxRows, nil)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			err:      nil,
		},
		{
			name: "UnsuccessfulReadProductsByPriceProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByRating, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			err:      ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsCategoryByRating(context.Background(), tc.id, tc.paging, tc.count, tc.ratingBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsCategoryByRatingPrice(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		ratingBy   string
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadCategoryByRatingPrice",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByRatingPrice, gomock.Any()).Return(pgxRows, nil)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			priceBy:  "%s",
			err:      nil,
		},
		{
			name: "UnsuccessfulReadCategoryByRatingPriceProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsCategoryByRatingPrice, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			priceBy:  "%s",
			err:      ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsCategoryByRatingPrice(context.Background(), tc.id, tc.paging, tc.count, tc.ratingBy, tc.priceBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsByRatingPrice(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		ratingBy   string
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadProductsByRatingPrice",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByRatingPrice, gomock.Any()).Return(pgxRows, nil)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			priceBy:  "%s",
			err:      nil,
		},
		{
			name: "UnsuccessfulReadProductsByRatingPriceProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByRatingPrice, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			priceBy:  "%s",
			err:      ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsByRatingPrice(context.Background(), tc.paging, tc.count, tc.ratingBy, tc.priceBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsByRating(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		ratingBy   string
		err        error
	}{
		{
			name: "SuccessfulReadProductsByRating",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByRating, gomock.Any()).Return(pgxRows, nil)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			err:      nil,
		},
		{
			name: "UnsuccessfulReadProductsByRatingProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByRating, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns:  []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:       2,
			paging:   4,
			count:    6,
			ratingBy: "%s",
			err:      ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsByRating(context.Background(), tc.paging, tc.count, tc.ratingBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProductsByPrice(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadProductsByPrice",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByPrice, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     nil,
		},
		{
			name: "UnsuccessfulReadProductsByPriceProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProductsByPrice, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProductsByPrice(context.Background(), tc.paging, tc.count, tc.priceBy)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProducts(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         int
		paging     int64
		count      int64
		priceBy    string
		err        error
	}{
		{
			name: "SuccessfulReadProducts",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     nil,
		},
		{
			name: "UnsuccessfulReadProductsProductNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getProducts, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "description", "price", "imgsrc", "rating", "category_id", "category_name", "count_comments"},
			id:      2,
			paging:  4,
			count:   6,
			priceBy: "%s",
			err:     ErrProductNotFound,
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProducts(context.Background(), tc.paging, tc.count)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProductsRepo_ReadProduct(t *testing.T) {
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

			repo := NewProductsRepo(mockPool)
			_, err := repo.ReadProduct(context.Background(), tc.id)

			assert.Equal(t, tc.err, err)
		})
	}
}

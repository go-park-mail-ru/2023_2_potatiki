package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryRepo_ReadCategories(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadCartProducts",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getCategories).Return(pgxRows, nil)
			},
			columns: []string{"id", "name", "count"},
			err:     nil,
		},
		{
			name: "UnsuccessfulReadCartProductsErrProductsNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getCategories).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "name", "count"},
			err:     ErrCategoryNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(int64(0), "", int64(0)).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewCategoryRepo(mockPool)
			_, err := repo.ReadCategories(context.Background())

			assert.Equal(t, tc.err, err)
		})
	}
}

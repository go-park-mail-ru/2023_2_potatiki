package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressRepo_ReadAllAddresses(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadAllAddresses",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), readAddresses, gomock.Any()).Return(pgxRows, nil)
			},
			columns: []string{"id", "city", "street", "house", "flat", "is_current"},
			err:     nil,
		},
		{
			name: "UnsuccessfulReadAllAddressesAddressesNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), readAddresses, gomock.Any()).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "city", "street", "house", "flat", "is_current"},
			err:     ErrAddressesNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", "", "", true).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewAddressRepo(mockPool)
			_, err := repo.ReadAllAddresses(context.Background(), userID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_CreateAddress(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulCreateAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), createAddress, gomock.Any(), gomock.Any(), gomock.Any(),
					gomock.Any(), gomock.Any(), gomock.Any(), true).Return(nil, nil)
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

			repo := NewAddressRepo(mockPool)
			_, err := repo.CreateAddress(context.Background(), userID, models.AddressPayload{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_DeleteAddress(t *testing.T) {
	testCases := []struct {
		name         string
		mockRepoFn   func(*pgxpoolmock.MockPgxPool)
		columns      []string
		err          error
		rowsAffected int64
	}{
		{
			name: "UnsuccessfulDeleteAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), deleteAddress, gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			columns:      []string{"id"},
			err:          ErrNoCurrentAddressNotFound,
			rowsAffected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewAddressRepo(mockPool)
			err := repo.DeleteAddress(context.Background(), models.AddressDelete{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_MakeCurrentAddress(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "UnsuccessfulMakeCurrentAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), makeCurrentAddress, gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			err: ErrCurrentAddressNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			tc.mockRepoFn(mockPool)

			repo := NewAddressRepo(mockPool)
			err := repo.MakeCurrentAddress(context.Background(), models.AddressMakeCurrent{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_ReadAddress(t *testing.T) {
	userID := uuid.NewV4()
	addressID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), readAddress, addressID, userID).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"city", "street", "house", "flat", "is_current"},
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow("", "", "", "", true).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewAddressRepo(mockPool)
			_, err := repo.ReadAddress(context.Background(), userID, addressID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_ReadCurrentAddress(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadCurrentAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), readCurrentAddress, userID).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"id", "city", "street", "house", "flat", "is_current"},
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, "", "", "", "", true).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewAddressRepo(mockPool)
			_, err := repo.ReadCurrentAddress(context.Background(), userID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_ReadCurrentAddressID(t *testing.T) {
	userID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadCurrentAddressID",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), readCurrentAddressID, gomock.Any()).Return(pgxRows)
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

			repo := NewAddressRepo(mockPool)
			_, err := repo.ReadCurrentAddressID(context.Background(), userID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAddressRepo_UpdateAddress(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulUpdateAddress",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), updateAddress, gomock.Any(), gomock.Any(), gomock.Any(),
					gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
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

			repo := NewAddressRepo(mockPool)
			err := repo.UpdateAddress(context.Background(), models.Address{})

			assert.Equal(t, tc.err, err)
		})
	}
}

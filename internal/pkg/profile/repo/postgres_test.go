package repo

import (
	"context"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProfileRepo_CreateProfile(t *testing.T) {
	t.Skip()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulCreateProfile",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), addProfile, gomock.Any(), gomock.Any(),
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

			repo := NewProfileRepo(mockPool)
			err := repo.CreateProfile(context.Background(), &models.Profile{})

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProfileRepo_GetProfileIdByLogin(t *testing.T) {
	login := "defaultLogin11"
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulGetProfileIdByLogin",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), profileIdByLogin, login).Return(pgxRows)
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

			repo := NewProfileRepo(mockPool)
			_, err := repo.GetProfileIdByLogin(context.Background(), login)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProfileRepo_ReadProfile(t *testing.T) {
	profileID := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		err        error
	}{
		{
			name: "SuccessfulReadProfile",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), profileById, profileID).Return(pgxRows)
				pgxRows.Next()
			},
			columns: []string{"login", "description", "imgsrc", "phone", "passwordhash"},
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow("", "", "", "", []byte{}).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewProfileRepo(mockPool)
			_, err := repo.ReadProfile(context.Background(), profileID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProfileRepo_UpdatePhoto(t *testing.T) {
	userID := uuid.NewV4()
	photoName := "default.png"
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulUpdatePhoto",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), updateProfilePhoto, photoName, userID).
					Return(nil, nil)
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

			repo := NewProfileRepo(mockPool)
			err := repo.UpdatePhoto(context.Background(), userID, photoName)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestProfileRepo_UpdateProfile(t *testing.T) {
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool)
		err        error
	}{
		{
			name: "SuccessfulUpdateProfile",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), updateProfileInfo, gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, nil)
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

			repo := NewProfileRepo(mockPool)
			err := repo.UpdateProfile(context.Background(), &models.Profile{})

			assert.Equal(t, tc.err, err)
		})
	}
}

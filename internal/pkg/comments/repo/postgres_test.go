package repo

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestCommentsRepo_ReadProductComments(t *testing.T) {
	uuidForTest := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		id         uuid.UUID
		err        error
	}{
		{
			name: "SuccessfulReadProductComments",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getComments, uuidForTest).Return(pgxRows, nil)
			},
			columns: []string{"id", "product_id", "pros", "cons", "comment", "rating", "user_name", "creation_date"},
			id:      uuidForTest,
			err:     nil,
		},
		{
			name: "UnsuccessfulReadProductCommentsErrCommentNotFound",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().Query(gomock.Any(), getComments, uuidForTest).Return(pgxRows, pgx.ErrNoRows)
			},
			columns: []string{"id", "product_id", "pros", "cons", "comment", "rating", "user_name", "creation_date"},
			id:      uuidForTest,
			err:     ErrCommentNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(uuid.UUID{}, uuid.UUID{}, "", "", "", 0, "", time.Time{}).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewCommentsRepo(mockPool)
			_, err := repo.ReadProductComments(context.Background(), tc.id)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCommentsRepo_ReadCountOfCommentsToProduct(t *testing.T) {
	uuidForTest := uuid.NewV4()
	testCases := []struct {
		name       string
		mockRepoFn func(*pgxpoolmock.MockPgxPool, pgx.Rows)
		columns    []string
		userID     uuid.UUID
		productID  uuid.UUID
		err        error
	}{
		{
			name: "SuccessfulReadCountOfCommentsToProduct",
			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
				mockPool.EXPECT().QueryRow(gomock.Any(), countOfCommentsToProduct, uuidForTest).Return(pgxRows)
				pgxRows.Next()
			},
			columns:   []string{"count", "id", "product_id", "pros", "cons", "comment", "rating", "user_name", "creation_date"},
			userID:    uuidForTest,
			productID: uuidForTest,
			err:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
			defer ctr.Finish()

			pgxRows := pgxpoolmock.NewRows(tc.columns).
				AddRow(0, uuid.UUID{}, uuid.UUID{}, "", "", "", 0, "", time.Time{}).ToPgxRows()

			tc.mockRepoFn(mockPool, pgxRows)

			repo := NewCommentsRepo(mockPool)
			_, _, err := repo.ReadCountOfCommentsToProduct(context.Background(), tc.userID, tc.productID)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCommentsRepo_MakeComment(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx            context.Context
		commentPayload models.CommentPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CommentsRepo{
				db: tt.fields.db,
			}
			got, err := r.MakeComment(tt.args.ctx, tt.args.commentPayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

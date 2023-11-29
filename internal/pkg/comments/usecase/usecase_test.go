package usecase

import (
	"context"
	"errors"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/comments/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/comments/repo"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestCommentsUsecase_CreateComment(t *testing.T) {
	type fields struct {
		repo *mock.MockCommentsRepo
	}
	type args struct {
		ctx            context.Context
		commentPayload models.CommentPayload
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CommentsUsecase
		args    args
		want    models.Comment
		wantErr error
	}{
		{
			name: "TestCommentsUsecase_ReadCountOfCommentsToProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCountOfCommentsToProduct(gomock.Any(),
					uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
					uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				).Return(1, models.Comment{}, nil)
			},
			args: args{context.Background(), models.CommentPayload{UserID: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				ProductID: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Comment{},
			wantErr: ErrManyCommentsToProduct,
		},
		{
			name: "TestCommentsUsecase_ReadCountOfCommentsToProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCountOfCommentsToProduct(gomock.Any(),
					uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000 "),
					uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000"),
				).Return(0, models.Comment{}, nil)
				f.repo.EXPECT().MakeComment(gomock.Any(),
					models.CommentPayload{},
				).Return(models.Comment{}, nil)
			},
			args: args{context.Background(), models.CommentPayload{UserID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000 "),
				ProductID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000 a")}},
			want:    models.Comment{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCommentsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCommentsUsecase(f.repo)

			_, err := tt.uc.CreateComment(tt.args.ctx, tt.args.commentPayload)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CartUsecase.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCommentsUsecase_GetProductComments(t *testing.T) {
	type fields struct {
		repo *mock.MockCommentsRepo
	}
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CommentsUsecase
		args    args
		want    []models.Comment
		wantErr error
	}{
		{
			name: "TestCommentsUsecase_ReadCountOfCommentsToProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProductComments(gomock.Any(),
					uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				).Return([]models.Comment{}, nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    []models.Comment{},
			wantErr: nil,
		},
		{
			name: "TestCommentsUsecase_ReadCountOfCommentsToProduct bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProductComments(gomock.Any(),
					uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				).Return([]models.Comment{}, repo.ErrCommentNotFound)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    []models.Comment{},
			wantErr: repo.ErrCommentNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCommentsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCommentsUsecase(f.repo)

			_, err := tt.uc.GetProductComments(tt.args.ctx, tt.args.productID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CartUsecase.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

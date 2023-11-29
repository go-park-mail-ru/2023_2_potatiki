package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestCommentsRepo_ReadCountOfCommentsToProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    uuid.UUID
		productID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CommentsRepo
		args    args
		want    int
		want1   models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.r.ReadCountOfCommentsToProduct(tt.args.ctx, tt.args.userID, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentsRepo.ReadCountOfCommentsToProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommentsRepo.ReadCountOfCommentsToProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CommentsRepo.ReadCountOfCommentsToProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCommentsRepo_MakeComment(t *testing.T) {
	type args struct {
		ctx            context.Context
		commentPayload models.CommentPayload
	}
	tests := []struct {
		name    string
		r       *CommentsRepo
		args    args
		want    models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MakeComment(tt.args.ctx, tt.args.commentPayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentsRepo.MakeComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentsRepo.MakeComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentsRepo_ReadProductComments(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CommentsRepo
		args    args
		want    []models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadProductComments(tt.args.ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentsRepo.ReadProductComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentsRepo.ReadProductComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

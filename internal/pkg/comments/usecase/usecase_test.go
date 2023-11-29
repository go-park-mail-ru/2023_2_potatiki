package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestCommentsUsecase_CreateComment(t *testing.T) {
	type args struct {
		ctx            context.Context
		commentPayload models.CommentPayload
	}
	tests := []struct {
		name    string
		uc      *CommentsUsecase
		args    args
		want    models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CreateComment(tt.args.ctx, tt.args.commentPayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentsUsecase.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentsUsecase.CreateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentsUsecase_GetProductComments(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}
	tests := []struct {
		name    string
		uc      *CommentsUsecase
		args    args
		want    []models.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetProductComments(tt.args.ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentsUsecase.GetProductComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentsUsecase.GetProductComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

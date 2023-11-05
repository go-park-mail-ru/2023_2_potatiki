package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category"
)

func TestCategoryUsecase_Categories(t *testing.T) {
	type fields struct {
		repo category.CategoryRepo
	}
	type args struct {
		ctx context.Context //nolint:containedctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CategoryTree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CategoryUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.Categories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryUsecase.Categories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CategoryUsecase.Categories() = %v, want %v", got, tt.want)
			}
		})
	}
}

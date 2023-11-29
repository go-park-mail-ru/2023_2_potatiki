package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

func TestCategoryRepo_ReadCategories(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *CategoryRepo
		args    args
		want    models.CategoryTree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadCategories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryRepo.ReadCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CategoryRepo.ReadCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

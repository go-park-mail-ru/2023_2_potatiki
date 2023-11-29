package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

func TestSearchUsecase_SearchProducts(t *testing.T) {
	type args struct {
		ctx         context.Context
		productName string
	}
	tests := []struct {
		name    string
		uc      *SearchUsecase
		args    args
		want    []models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.SearchProducts(tt.args.ctx, tt.args.productName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchUsecase.SearchProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUsecase.SearchProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

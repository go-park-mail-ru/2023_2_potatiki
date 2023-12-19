package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestPromoUsecase_CheckPromocode(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
		name   string
	}
	tests := []struct {
		name    string
		uc      *PromoUsecase
		args    args
		want    *models.Promocode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CheckPromocode(tt.args.ctx, tt.args.userID, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromoUsecase.CheckPromocode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PromoUsecase.CheckPromocode() = %v, want %v", got, tt.want)
			}
		})
	}
}

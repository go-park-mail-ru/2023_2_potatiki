package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

func TestPromoRepo_ReadPromocode(t *testing.T) {
	type args struct {
		ctx           context.Context
		promocodeName string
	}
	tests := []struct {
		name    string
		r       *PromoRepo
		args    args
		want    *models.Promocode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadPromocode(tt.args.ctx, tt.args.promocodeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromoRepo.ReadPromocode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PromoRepo.ReadPromocode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category/mocks"

	"github.com/golang/mock/gomock"
)

func TestCategoryUsecase_Categories(t *testing.T) {
	type fields struct {
		repo *mock.MockCategoryRepo
	}
	type args struct {
		ctx context.Context //nolint:containedctx
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CategoryUsecase
		args    args
		want    models.CategoryTree
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_UpdateData good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCategories(gomock.Any()).Return(models.CategoryTree{}, nil)
			},
			args:    args{context.Background()},
			want:    models.CategoryTree{},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_UpdateData bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCategories(gomock.Any()).Return(models.CategoryTree{}, errors.New("dummyError"))
			},
			args:    args{context.Background()},
			want:    models.CategoryTree{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCategoryRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCategoryUsecase(f.repo)

			got, err := tt.uc.Categories(tt.args.ctx)
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

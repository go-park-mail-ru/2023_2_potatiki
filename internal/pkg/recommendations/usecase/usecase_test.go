package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations/repo"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestRecommendationsUsecase_AnonRecommendations(t *testing.T) {
	type fields struct {
		repo *mock.MockRecommendationsRepo
	}
	type args struct {
		ctx        context.Context
		productID  uuid.UUID
		categoryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(f *fields)
		uc      *RecommendationsUsecase
		want    models.ProductSlice
		wantErr bool
	}{
		{
			name: "SuccessfullyAnonRecommendations",
			args: args{
				context.Background(),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				6,
			},
			want:    models.ProductSlice{},
			wantErr: false,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProductsFromParentCategory(
					gomock.Any(),
					int64(6),
				).Return(models.ProductSlice{}, nil)
			},
		},
		{
			name: "FailureAnonRecommendations",
			args: args{
				context.Background(),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				6,
			},
			want:    models.ProductSlice{},
			wantErr: true,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProductsFromParentCategory(
					gomock.Any(),
					int64(6),
				).Return(models.ProductSlice{}, errors.New("error read products"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockRecommendationsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewRecommendationsUsecase(f.repo)

			got, err := tt.uc.AnonRecommendations(tt.args.ctx, tt.args.productID, tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecommendationsUsecase.AnonRecommendations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecommendationsUsecase.AnonRecommendations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecommendationsUsecase_Recommendations(t *testing.T) {
	type fields struct {
		repo *mock.MockRecommendationsRepo
	}
	type args struct {
		ctx        context.Context
		productID  uuid.UUID
		userID     uuid.UUID
		categoryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(f *fields)
		uc      *RecommendationsUsecase
		want    models.ProductSlice
		wantErr bool
	}{
		{
			name: "RecommendationsWithZeroActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				6,
			},
			want:    models.ProductSlice{},
			wantErr: false,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, nil)
				f.repo.EXPECT().ReadProductsFromParentCategory(
					gomock.Any(),
					int64(6),
				).Return(models.ProductSlice{}, nil)
			},
		},
		{
			name: "RecommendationsWithActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				6,
			},
			want:    models.ProductSlice{},
			wantErr: false,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{
					Product: models.ProductStatisticMap{
						uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"): models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: 2,
						},
					},
					Category: models.CategoryStatisticMap{
						1: models.CategoryStatistic{
							CategoryID:     int64(1),
							ActivityPoints: 3,
						},
					},
				}, nil)
				f.repo.EXPECT().ReadProduct(
					gomock.Any(),
					gomock.Any(),
				).Return(models.Product{}, nil)
				f.repo.EXPECT().ReadProductsFromCategory(
					gomock.Any(),
					gomock.Any(),
				).Return(models.ProductSlice{}, nil)
				f.repo.EXPECT().ReadProductsFromParentCategory(
					gomock.Any(),
					int64(6),
				).Return(models.ProductSlice{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockRecommendationsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewRecommendationsUsecase(f.repo)

			_, err := tt.uc.Recommendations(tt.args.ctx, tt.args.userID, tt.args.productID, tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecommendationsUsecase.Recommendations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRecommendationsUsecase_UpdateUserActivity(t *testing.T) {
	type fields struct {
		repo *mock.MockRecommendationsRepo
	}
	type args struct {
		ctx         context.Context
		userID      uuid.UUID
		newActivity models.UserActivity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(f *fields)
		uc      *RecommendationsUsecase
		wantErr bool
	}{
		{
			name: "SuccessfullyUpdateUserActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				models.UserActivity{
					Product: models.ProductStatisticSlice{
						models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: int64(2),
						},
					},
					Category: models.CategoryStatisticSlice{
						models.CategoryStatistic{
							CategoryID:     int64(6),
							ActivityPoints: int64(2),
						},
					},
				},
			},
			wantErr: false,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, nil)
				f.repo.EXPECT().UpdateUserActivity(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(nil)
			},
		},
		{
			name: "FailureUpdateUserActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				models.UserActivity{
					Product: models.ProductStatisticSlice{
						models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: int64(2),
						},
					},
					Category: models.CategoryStatisticSlice{
						models.CategoryStatistic{
							CategoryID:     int64(6),
							ActivityPoints: int64(2),
						},
					},
				},
			},
			wantErr: true,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, nil)
				f.repo.EXPECT().UpdateUserActivity(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(errors.New("error update user activity"))
			},
		},
		{
			name: "FailureReadUserActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				models.UserActivity{
					Product: models.ProductStatisticSlice{
						models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: int64(2),
						},
					},
					Category: models.CategoryStatisticSlice{
						models.CategoryStatistic{
							CategoryID:     int64(6),
							ActivityPoints: int64(2),
						},
					},
				},
			},
			wantErr: true,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, errors.New("error read activity"))
			},
		},
		{
			name: "SuccessfullyCreateUpdateUserActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				models.UserActivity{
					Product: models.ProductStatisticSlice{
						models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: int64(2),
						},
					},
					Category: models.CategoryStatisticSlice{
						models.CategoryStatistic{
							CategoryID:     int64(6),
							ActivityPoints: int64(2),
						},
					},
				},
			},
			wantErr: false,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, repo.ErrActivityNotFound)
				f.repo.EXPECT().CreateUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(nil)
				f.repo.EXPECT().UpdateUserActivity(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(nil)
			},
		},
		{
			name: "FailureCreateUpdateUserActivity",
			args: args{
				context.WithValue(context.Background(), "zuzu-t", uuid.NewV4()),
				uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
				models.UserActivity{
					Product: models.ProductStatisticSlice{
						models.ProductStatistic{
							ProductID:      uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
							ActivityPoints: int64(2),
						},
					},
					Category: models.CategoryStatisticSlice{
						models.CategoryStatistic{
							CategoryID:     int64(6),
							ActivityPoints: int64(2),
						},
					},
				},
			},
			wantErr: true,
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(models.UserActivityStore{}, repo.ErrActivityNotFound)
				f.repo.EXPECT().CreateUserActivity(
					gomock.Any(),
					gomock.Any(),
				).Return(errors.New("err create user activity"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockRecommendationsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewRecommendationsUsecase(f.repo)

			err := tt.uc.UpdateUserActivity(tt.args.ctx, tt.args.userID, tt.args.newActivity)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecommendationsUsecase.UpdateUserActivity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

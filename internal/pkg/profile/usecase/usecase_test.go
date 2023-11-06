package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/mocks"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestProfileUsecase_GetProfile(t *testing.T) {
	type fields struct {
		repo *mock.MockProfileRepo
		cfg  *mock.MockProfileConfig
	}
	type args struct {
		ctx context.Context //nolint:containedctx
		Id  uuid.UUID
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProfileUsecase
		args    args
		want    *models.Profile
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_GetProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{Login: "s"}, nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    &models.Profile{Login: "s"},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_GetProduct bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, fmt.Errorf("dummyErr"))
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProfileRepo(ctrl),
				cfg:  mock.NewMockProfileConfig(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProfileUsecase(f.repo, f.cfg)

			got, err := tt.uc.GetProfile(tt.args.ctx, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileUsecase.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileUsecase.GetProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileUsecase_UpdateData(t *testing.T) {
	type fields struct {
		repo *mock.MockProfileRepo
		cfg  *mock.MockProfileConfig
	}
	type args struct {
		ctx     context.Context //nolint:containedctx
		Id      uuid.UUID
		payload *models.UpdateProfileDataPayload
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProfileUsecase
		args    args
		want    *models.Profile
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_UpdateData good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, nil)
				f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{Phone: "+12345678"}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Phone: "+12345678"}},
			want:    &models.Profile{Phone: "+12345678"},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_UpdateData bad read repo",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, fmt.Errorf("dummyErr"))
				//f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Phone: "+12345678"}},
			wantErr: true,
		},
		{
			name: "TestProductsUsecase_UpdateData bad update repo",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, nil)
				f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{Phone: "+12345678"}).Return(fmt.Errorf("dummyErr"))
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Phone: "+12345678"}},
			wantErr: true,
		},
		{
			name: "TestProductsUsecase_UpdateData bad validation",
			prepare: func(f *fields) {
				//f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, fmt.Errorf("dummyErr"))
				//f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{}},
			wantErr: true,
		},
		{
			name: "TestProductsUsecase_UpdateData bad validation",
			prepare: func(f *fields) {
				//f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, fmt.Errorf("dummyErr"))
				//f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args: args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Passwords: struct {
				OldPass string "json:\"oldPass\" validate:\"omitempty,min=8,max=32\""
				NewPass string "json:\"newPass\" validate:\"omitempty,min=8,max=32\""
			}{OldPass: "123", NewPass: "123"}}},
			wantErr: true,
		},
		{
			name: "TestProductsUsecase_UpdateData bad CheckPass",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{PasswordHash: []byte("12345678")}, nil)
				//f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args: args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Passwords: struct {
				OldPass string "json:\"oldPass\" validate:\"omitempty,min=8,max=32\""
				NewPass string "json:\"newPass\" validate:\"omitempty,min=8,max=32\""
			}{OldPass: "12345678", NewPass: "12345678"}}},
			wantErr: true,
		},
		/*{
			name: "TestProductsUsecase_UpdateData good CheckPass",
			prepare: func(f *fields) {
				pass := hasher.HashPass("12345678")
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{PasswordHash: pass}, nil)
				f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{PasswordHash: pass}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args: args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{Passwords: struct {
				OldPass string "json:\"oldPass\" validate:\"omitempty,min=8,max=32\""
				NewPass string "json:\"newPass\" validate:\"omitempty,min=8,max=32\""
			}{OldPass: "12345678", NewPass: "12345678"}}},
			wantErr: false,
		},*/
		{
			name: "TestProductsUsecase_UpdateData bad validation",
			prepare: func(f *fields) {
				//f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{}, fmt.Errorf("dummyErr"))
				//f.repo.EXPECT().UpdateProfile(gomock.Any(), &models.Profile{}).Return(nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args: args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), &models.UpdateProfileDataPayload{
				Phone: "+1234567890",
				Passwords: struct {
					OldPass string "json:\"oldPass\" validate:\"omitempty,min=8,max=32\""
					NewPass string "json:\"newPass\" validate:\"omitempty,min=8,max=32\""
				}{OldPass: "12345678", NewPass: ""}}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProfileRepo(ctrl),
				cfg:  mock.NewMockProfileConfig(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProfileUsecase(f.repo, f.cfg)

			got, err := tt.uc.UpdateData(tt.args.ctx, tt.args.Id, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileUsecase.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileUsecase.UpdateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileUsecase_UpdatePhoto(t *testing.T) {
	type fields struct {
		repo *mock.MockProfileRepo
		cfg  *mock.MockProfileConfig
	}
	type args struct {
		ctx           context.Context //nolint:containedctx
		Id            uuid.UUID
		filePhotoByte []byte
		fileType      string
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProfileUsecase
		args    args
		want    *models.Profile
		wantErr bool
	}{
		/*{
			name: "TestProductsUsecase_UpdateData good",
			prepare: func(f *fields) {
				m := fstest.MapFS{
					"photos/hello.png": {
						Data: []byte("hello, world"),
					},
				}
				m.Stat("photos/")
				f.repo.EXPECT().UpdatePhoto(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), "image/png").Return(nil)
				f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{ImgSrc: "image/png"}, nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("photos/")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), []byte{1, 2, 3}, "image/png"},
			want:    &models.Profile{ImgSrc: "image/png"},
			wantErr: false,
		},*/
		{
			name: "TestProductsUsecase_UpdateData bad",
			prepare: func(f *fields) {
				//f.repo.EXPECT().UpdatePhoto(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), "image/png").Return(nil)
				//f.repo.EXPECT().ReadProfile(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(&models.Profile{ImgSrc: "image/png"}, nil)
				f.cfg.EXPECT().GetPhotosFilePath().Return("")
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), []byte{1, 2, 3}, "image/png"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProfileRepo(ctrl),
				cfg:  mock.NewMockProfileConfig(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProfileUsecase(f.repo, f.cfg)

			got, err := tt.uc.UpdatePhoto(tt.args.ctx, tt.args.Id, tt.args.filePhotoByte, tt.args.fileType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileUsecase.UpdatePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileUsecase.UpdatePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

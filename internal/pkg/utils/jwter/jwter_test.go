package jwter

import (
	"reflect"
	"testing"
	"time"

	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func Test_jwtManager_DecodeAuthToken(t *testing.T) {
	type fields struct {
		cfg *mock.MockConfiger
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *jwtManager
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "Test_jwtManager_DecodeAuthToken good",
			prepare: func(f *fields) {
				f.cfg.EXPECT().GetIssuer().Return("auth")
				f.cfg.EXPECT().GetSecret().Return("your-256-bit-secret")
				f.cfg.EXPECT().GetTTL().Return(time.Hour * 6)
			},
			args:    args{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			want:    uuid.UUID{},
			wantErr: false,
		},
		{
			name: "Test_jwtManager_DecodeAuthToken bad",
			prepare: func(f *fields) {
				f.cfg.EXPECT().GetIssuer().Return("auth")
				f.cfg.EXPECT().GetSecret().Return("your-256-bit-secret")
				f.cfg.EXPECT().GetTTL().Return(time.Hour * 6)
			},
			args:    args{"-------flKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			want:    uuid.UUID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				cfg: mock.NewMockConfiger(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = New(f.cfg)

			got, err := tt.uc.DecodeAuthToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileUsecase.DecodeAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileUsecase.DecodeAuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtManager_EncodeCSRFToken(t *testing.T) {
	type fields struct {
		cfg *mock.MockConfiger
	}
	type args struct {
		UserAgent string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *jwtManager
		args    args
		want    string
		want1   time.Time
		wantErr bool
	}{ /*
			{
				name: "Test_jwtManager_EncodeCSRFToken good",
				prepare: func(f *fields) {
					f.cfg.EXPECT().GetIssuer().Return("auth")
					f.cfg.EXPECT().GetSecret().Return("your-256-bit-secret")
					f.cfg.EXPECT().GetTTL().Return(time.Hour * 6)
				},
				args:    args{"UserAgent"},
				want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyQWdlbnQiOiJVc2VyQWdlbnQiLCJpc3MiOiJhdXRoIiwiZXhwIjoxNjk5NDE2MzUyfQ.Wjf-vxjri1gQN638ZV0omxINoX9yag1AfB4Wunrz3Rg",
				wantErr: false,
			},
			{
				name: "Test_jwtManager_EncodeCSRFToken bad",
				prepare: func(f *fields) {
					f.cfg.EXPECT().GetIssuer().Return("auth")
					f.cfg.EXPECT().GetSecret().Return("your-256-bit-secret")
					f.cfg.EXPECT().GetTTL().Return(time.Hour * 6)
				},
				args:    args{"-------UserAgent"},
				want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyQWdlbnQiOiJVc2VyQWdlbnQiLCJpc3MiOiJhdXRoIiwiZXhwIjoxNjk5NDE2MDQ1fQ.EE_s_O_azore9NYDhoqqW_Z1kCPuxhRJOV3HBuzgIdI",
				want1:   time.Now(),
				wantErr: true,
			},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				cfg: mock.NewMockConfiger(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = New(f.cfg)

			got, got1, err := tt.uc.EncodeCSRFToken(tt.args.UserAgent)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtManager.EncodeCSRFToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("jwtManager.EncodeCSRFToken() got = %v, want %v", got, tt.want)
			}
			if !tt.wantErr && !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("jwtManager.EncodeCSRFToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtManager_DecodeCSRFToken(t *testing.T) {
	type fields struct {
		ttl    time.Duration
		secret string
		issuer string
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtManager{
				ttl:    tt.fields.ttl,
				secret: tt.fields.secret,
				issuer: tt.fields.issuer,
			}
			got, err := j.DecodeCSRFToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtManager.DecodeCSRFToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtManager.DecodeCSRFToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtManager_generateToken(t *testing.T) {
	type fields struct {
		ttl    time.Duration
		secret string
		issuer string
	}
	type args struct {
		claims jwt.Claims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtManager{
				ttl:    tt.fields.ttl,
				secret: tt.fields.secret,
				issuer: tt.fields.issuer,
			}
			got, got1, err := j.generateToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtManager.generateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtManager.generateToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("jwtManager.generateToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_jwtManager_getKeyFunc(t *testing.T) {
	type fields struct {
		ttl    time.Duration
		secret string
		issuer string
	}
	tests := []struct {
		name   string
		fields fields
		want   jwt.Keyfunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtManager{
				ttl:    tt.fields.ttl,
				secret: tt.fields.secret,
				issuer: tt.fields.issuer,
			}
			if got := j.getKeyFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jwtManager.getKeyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

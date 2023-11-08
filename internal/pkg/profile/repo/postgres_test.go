package repo

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestProfileRepo_CreateProfile(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx context.Context
		p   *models.Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepo{
				db: tt.fields.db,
			}
			if err := r.CreateProfile(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("CreateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileRepo_GetProfileIdByLogin(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepo{
				db: tt.fields.db,
			}
			got, err := r.GetProfileIdByLogin(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfileIdByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProfileIdByLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_ReadProfile(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx context.Context
		Id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepo{
				db: tt.fields.db,
			}
			got, err := r.ReadProfile(tt.args.ctx, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_UpdatePhoto(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx       context.Context
		userID    uuid.UUID
		photoName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepo{
				db: tt.fields.db,
			}
			if err := r.UpdatePhoto(tt.args.ctx, tt.args.userID, tt.args.photoName); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileRepo_UpdateProfile(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
	}
	type args struct {
		ctx context.Context
		p   *models.Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepo{
				db: tt.fields.db,
			}
			if err := r.UpdateProfile(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

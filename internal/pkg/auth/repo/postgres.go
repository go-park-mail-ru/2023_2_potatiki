package repo

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

type AuthRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(context.Context, models.User) (models.Profile, error) {
	panic("unimplemented")
}

func (r *AuthRepo) CheckUser(context.Context, models.User) (models.Profile, error) {
	panic("unimplemented")
}
func (r *AuthRepo) ReadProfile(context.Context, uuid.UUID) (models.Profile, error) {
	panic("unimplemented")
}

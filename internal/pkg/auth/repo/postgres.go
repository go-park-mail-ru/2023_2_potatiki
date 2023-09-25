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

func NewAuthRepo(NewDb *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: NewDb,
	}
}

func (r *AuthRepo) CreateUser(context.Context, models.User) (models.Profile, error) {
	return models.Profile{}, nil
}

func (r *AuthRepo) CheckUser(context.Context, models.User) (models.Profile, error) {
	return models.Profile{}, nil
}
func (r *AuthRepo) ReadProfile(context.Context, uuid.UUID) (models.Profile, error) {
	return models.Profile{}, nil
}

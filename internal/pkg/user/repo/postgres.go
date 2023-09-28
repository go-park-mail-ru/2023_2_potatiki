package repo

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) ReadProduct(context.Context, uuid.UUID) (models.Profile, error) {
	panic("unimplemented")
}

package repo

import (
	"context"
	"database/sql"
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

func (r *UserRepo) UpdatePhoto(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	return nil
}

func (r *UserRepo) UpdateDescription(ctx context.Context, userID uuid.UUID, description string) error {
	return nil
}

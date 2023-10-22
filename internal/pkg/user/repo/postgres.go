package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
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

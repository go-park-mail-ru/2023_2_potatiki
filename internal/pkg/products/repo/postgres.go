package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

var (
	ErrPoductNotFound = errors.New("product not found")
)

func (r *ProductsRepo) ReadProduct(context.Context, uuid.UUID) (models.Product, error) {
	panic("unimplemented")
}

func (r *ProductsRepo) ReadProducts(context.Context, int64, int64) ([]models.Product, error) {
	panic("unimplemented")
}

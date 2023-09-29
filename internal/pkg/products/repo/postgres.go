package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

const (
	getProduct = "SELECT * FROM public.products WHERE id=$1;"
	//addProfile    = "INSERT INTO public.profiles(Id, Login, Description, ImgSrc, PasswordHash) VALUES($1, $2, $3, $4, $5);"
)

var (
	ErrPoductNotFound = errors.New("product not found")
)

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (r *ProductsRepo) ReadProduct(ctx context.Context, id uuid.UUID) (models.Product, error) {
	pr := models.Product{}
	err := r.db.QueryRowContext(ctx, getProduct, id).Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, ErrPoductNotFound
		}
		return models.Product{}, err
	}
	return pr, nil
}

func (r *ProductsRepo) ReadProducts(context.Context, int64, int64) ([]models.Product, error) {
	panic("unimplemented")
}

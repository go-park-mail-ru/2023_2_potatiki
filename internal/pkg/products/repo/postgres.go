package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

const (
	getProduct  = "SELECT * FROM products WHERE id=$1;"
	getProducts = "SELECT Id , NameProduct, Description, Price FROM products ORDER BY id LIMIT $1 OFFSET $2"
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

func (r *ProductsRepo) ReadProducts(ctx context.Context, offset int64, pageSize int64) ([]models.Product, error) {
	var productSlice []models.Product
	rows, err := r.db.QueryContext(ctx, getProducts, pageSize, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Product{}, ErrPoductNotFound
		}
		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price); err != nil {
			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)

	return productSlice, nil
}

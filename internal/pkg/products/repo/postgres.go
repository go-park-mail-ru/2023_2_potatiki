package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
)

const (
	getProduct  = "SELECT id, name, description, price, imgsrc, rating FROM product WHERE id=$1;"
	getProducts = `SELECT id, name, description, price, imgsrc, rating 
	FROM product ORDER BY id LIMIT $1 OFFSET $2;`
	getProductsByategoryId = `SELECT id, name, description, price, imgsrc, rating 
	FROM product WHERE category_id=$3 ORDER BY id LIMIT $1 OFFSET $2;`
)

var (
	ErrPoductNotFound = errors.New("product not found")
)

type ProductsRepo struct {
	db pgxtype.Querier // TODO: add logger
}

func NewProductsRepo(db pgxtype.Querier) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (r *ProductsRepo) ReadProduct(ctx context.Context, id uuid.UUID) (models.Product, error) {
	pr := models.Product{}
	err := r.db.QueryRow(ctx, getProduct, id).
		Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price, &pr.ImgSrc, &pr.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Product{}, err
	}

	return pr, nil
}

func (r *ProductsRepo) ReadProducts(ctx context.Context, paging int64, count int64) ([]models.Product, error) {
	var productSlice []models.Product
	rows, err := r.db.Query(ctx, getProducts, count, paging)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.ImgSrc, &product.Rating)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer rows.Close()

	return productSlice, nil
}

func (r *ProductsRepo) ReadCategory(ctx context.Context, id int, paging, count int64) ([]models.Product, error) {
	var productSlice []models.Product
	rows, err := r.db.Query(ctx, getProductsByategoryId, count, paging, id)
	fmt.Println(count, paging, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.ImgSrc, &product.Rating)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer rows.Close()

	return productSlice, nil
}

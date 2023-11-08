package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
)

const (
	getProduct = `SELECT p.id, p.name, p.description, p.price, p.imgsrc, p.rating, p.category_id,
    c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	WHERE p.id = $1;`

	getProducts = `SELECT p.id, p.name, p.description, p.price, p.imgsrc, p.rating, p.category_id,
    c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	ORDER BY p.id
	LIMIT $1 OFFSET $2;`

	getProductsByCategoryID = `SELECT p.id, p.name, p.description, p.price, p.imgsrc, p.rating, p.category_id,
    c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	WHERE p.category_id = $3
	ORDER BY p.id
	LIMIT $1 OFFSET $2;`
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
		Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price, &pr.ImgSrc, &pr.Rating, &pr.Category.Id, &pr.Category.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Product{}, err
	}

	return pr, nil
}

func (r *ProductsRepo) ReadProducts(ctx context.Context, paging int64, count int64) ([]models.Product, error) {
	productSlice := make([]models.Product, 0)
	rows, err := r.db.Query(ctx, getProducts, count, paging)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Category.Id,
			&product.Category.Name,
		)
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
	productSlice := make([]models.Product, 0)
	rows, err := r.db.Query(ctx, getProductsByCategoryID, count, paging, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Category.Id,
			&product.Category.Name,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer rows.Close()

	return productSlice, nil
}

package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
)

const (
	getProductsByName = `
	SELECT p.id,
	   p.name,
	   p.description,
	   p.price,
	   p.imgsrc,
	   p.rating,
	   p.category_id,
	   c.name AS category_name
	FROM product p
		 JOIN category c ON p.category_id = c.id
	WHERE similarity(c.name, $1) > 0.2 OR similarity(p.name, $1) > 0.1 OR similarity(p.description, $1) > 0.04
	ORDER BY similarity(c.name, $1) DESC, similarity(p.name, $1) DESC, similarity(p.description, $1) DESC
	LIMIT 10;`
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type SearchRepo struct {
	db pgxtype.Querier // TODO: add logger
}

func NewSearchRepo(db pgxtype.Querier) *SearchRepo {
	return &SearchRepo{
		db: db,
	}
}

func (r *SearchRepo) ReadProductsByName(ctx context.Context, productName string) ([]models.Product, error) {

	count := 10
	productSlice := make([]models.Product, 0, count)
	rows, err := r.db.Query(ctx, getProductsByName, productName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrProductNotFound
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

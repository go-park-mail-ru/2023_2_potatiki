package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	getProduct = `SELECT p.id, p.name, p.description, p.price, p.imgsrc,  p.category_id,
   c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	WHERE p.id = $1;`

	getProductsASCRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price ASC, COALESCE(AVG(cm.rating), 0) ASC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsASCRatingDESCPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	COALESCE(AVG(cm.rating), 0) ASC, p.price DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsDESCRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price DESC, COALESCE(AVG(cm.rating), 0) DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsDESCRatingASCPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price ASC, COALESCE(AVG(cm.rating), 0) DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	//getProductsByCategoryID = `SELECT p.id, p.name, p.description, p.price, p.imgsrc,  p.category_id,
	//c.name AS category_name
	//FROM product p
	//JOIN category c ON p.category_id = c.id
	//WHERE p.category_id = $3
	//ORDER BY p.id
	//LIMIT $1 OFFSET $2;`

	getProductsByCategoryIDASCRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price ASC, COALESCE(AVG(cm.rating), 0) ASC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByCategoryIDASCRatingDESCPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	COALESCE(AVG(cm.rating), 0) ASC, p.price DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByCategoryIDDESCRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price DESC, COALESCE(AVG(cm.rating), 0) DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByCategoryIDDESCRatingASCPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name
	ORDER BY
   	p.price ASC, COALESCE(AVG(cm.rating), 0) DESC
	LIMIT
		$1
	OFFSET
		$2;
	`
)

var (
	ErrProductNotFound = errors.New("product not found")
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
		Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price, &pr.ImgSrc, &pr.Category.Id, &pr.Category.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Product{}, ErrProductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Product{}, err
	}

	return pr, nil
}

func (r *ProductsRepo) ReadProducts(ctx context.Context, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0, count)
	var (
		rows pgx.Rows
		err  error
	)
	if ratingBy == "ASC" {
		if priceBy == "ASC" {
			rows, err = r.db.Query(ctx, getProductsASCRatingPrice, count, paging)
		} else {
			rows, err = r.db.Query(ctx, getProductsASCRatingDESCPrice, count, paging)
		}
	} else {
		if priceBy == "DESC" {
			rows, err = r.db.Query(ctx, getProductsDESCRatingPrice, count, paging)
		} else {
			rows, err = r.db.Query(ctx, getProductsDESCRatingASCPrice, count, paging)
		}
	}
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

func (r *ProductsRepo) ReadCategory(ctx context.Context, id int, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0)
	var (
		rows pgx.Rows
		err  error
	)
	if ratingBy == "ASC" {
		if priceBy == "ASC" {
			rows, err = r.db.Query(ctx, getProductsByCategoryIDASCRatingPrice, count, paging, id)
		} else {
			rows, err = r.db.Query(ctx, getProductsByCategoryIDASCRatingDESCPrice, count, paging, id)
		}
	} else {
		if priceBy == "DESC" {
			rows, err = r.db.Query(ctx, getProductsByCategoryIDDESCRatingPrice, count, paging, id)
		} else {
			rows, err = r.db.Query(ctx, getProductsByCategoryIDDESCRatingASCPrice, count, paging, id)
		}
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrProductNotFound
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

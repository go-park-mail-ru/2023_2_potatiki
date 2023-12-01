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
	getProduct = `SELECT p.id, p.name, p.description, p.price, p.imgsrc, COALESCE(AVG(cm.rating), 0) AS average_rating,
       p.category_id, c.name AS category_name, p.count_comments
	FROM product p
	JOIN category c ON p.category_id = c.id
	LEFT JOIN comment cm ON p.id = cm.productID
	WHERE p.id = $1
	GROUP BY p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments;
	`

	getProductsByRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		p.price %s, COALESCE(AVG(cm.rating), 0) %s
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		p.price %s
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByRating = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		COALESCE(AVG(cm.rating), 0) %s
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProducts = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		p.id DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsByCategoryID = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		p.creation_at DESC
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsCategoryByRatingPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		p.price %s, COALESCE(AVG(cm.rating), 0) %s
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsCategoryByRating = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
		COALESCE(AVG(cm.rating), 0) %s
	LIMIT
		$1
	OFFSET
		$2;
	`

	getProductsCategoryByPrice = `
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE
	    p.category_id = $3
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
		p.price %s
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
		Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price, &pr.ImgSrc, &pr.Rating,
			&pr.Category.Id, &pr.Category.Name, &pr.CountComments)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Product{}, ErrProductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Product{}, err
	}

	return pr, nil
}

func (r *ProductsRepo) ReadProducts(ctx context.Context, paging, count int64) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0, count)

	rows, err := r.db.Query(ctx, getProducts, count, paging)
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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsByPrice(ctx context.Context, paging, count int64, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0, count)

	query := fmt.Sprintf(getProductsByPrice, priceBy) //спросить про %S
	rows, err := r.db.Query(ctx, query, count, paging)

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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsByRating(ctx context.Context, paging, count int64, ratingBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0, count)

	query := fmt.Sprintf(getProductsByRating, ratingBy) //спросить про %S
	rows, err := r.db.Query(ctx, query, count, paging)

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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsByRatingPrice(ctx context.Context, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0, count)

	query := fmt.Sprintf(getProductsByRatingPrice, priceBy, ratingBy) //спросить про %S
	rows, err := r.db.Query(ctx, query, count, paging)
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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsCategoryByRatingPrice(ctx context.Context, id int, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0)

	query := fmt.Sprintf(getProductsCategoryByRatingPrice, priceBy, ratingBy)
	rows, err := r.db.Query(ctx, query, count, paging, id)
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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsCategoryByRating(ctx context.Context, id int, paging, count int64, ratingBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0)

	query := fmt.Sprintf(getProductsCategoryByRating, ratingBy)
	rows, err := r.db.Query(ctx, query, count, paging, id)
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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsCategoryByPrice(ctx context.Context, id int, paging, count int64, priceBy string) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0)

	query := fmt.Sprintf(getProductsCategoryByPrice, priceBy)
	rows, err := r.db.Query(ctx, query, count, paging, id)
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
			&product.CountComments,
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

func (r *ProductsRepo) ReadProductsCategory(ctx context.Context, id int, paging, count int64) (
	[]models.Product, error) {
	productSlice := make([]models.Product, 0)

	rows, err := r.db.Query(ctx, getProductsByCategoryID, count, paging, id)
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
			&product.CountComments,
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

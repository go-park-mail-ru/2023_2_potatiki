package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	getCart     = "SELECT id FROM cart WHERE Profile_id=$1 and is_current = true;"
	createCart  = "INSERT INTO cart(id, profile_id, is_current) VALUES($1, $2, true);"
	getProducts = "SELECT p._id , p.name_product, p.Description, p.Price, p.ImgSrc, p.Rating, sc.Quantity " +
		"FROM shopping_cart_item sc JOIN product p ON sc.product_id=p.id WHERE p.id=$1;"
	getProduct = "SELECT id FROM product where id=$1;"
	addProduct = "insert into shopping_cart_item(cart_id, product_id, quantity) values ($1, $2, $3)" +
		" ON CONFLICT ON CONSTRAINT uq_shopping_cart_item_cart_id_product_id " +
		"do update set quantity=$3 WHERE shopping_cart_item.cart_id=$1 and shopping_cart_item.product_id=$2;"
	deleteProduct = "DELETE FROM shopping_cart_item WHERE cart_id=$1 and product_id=$2;"
	deleteCard    = "DELETE FROM cart WHERE profile_id=$1"
)

var (
	ErrCartNotFound    = errors.New("cart not found")
	ErrPoductsNotFound = errors.New("products not found")
	ErrPoductNotFound  = errors.New("product not found")
)

type CartRepo struct {
	db pgxtype.Querier
}

func NewCartRepo(db pgxtype.Querier) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (r *CartRepo) CreateCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
	cartID := uuid.New()
	_, err := r.db.Exec(ctx, createCart, cartID, userID)
	if err != nil {
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return models.Cart{}, err
	}
	cart := models.Cart{}
	cart.Id = cartID
	cart.ProfileId = userID
	return cart, nil
}

func (r *CartRepo) DeleteCart(ctx context.Context, cartID uuid.UUID) error {
	_, err := r.db.Exec(ctx, deleteCard, cartID)
	if err != nil {
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return err
	}

	return nil
}

func (r *CartRepo) CheckCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
	cart := models.Cart{}
	err := r.db.QueryRow(ctx, getCart, userID).Scan(&cart.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Cart{}, ErrCartNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Cart{}, err
	}
	return cart, nil
}

func (r *CartRepo) ReadCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
	cart, err := r.CheckCart(ctx, userID)
	if err != nil {
		return models.Cart{}, err
	}

	cart, err = r.ReadCartProducts(ctx, cart)

	return cart, err
}

func (r *CartRepo) ReadCartProducts(ctx context.Context, cart models.Cart) (models.Cart, error) {
	rows, err := r.db.Query(ctx, getProducts, cart.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cart, ErrPoductsNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return cart, err
	}

	product := models.CartProduct{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Quantity,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return cart, err
		}
		cart.Products = append(cart.Products, product)
	}
	defer rows.Close()

	return cart, nil
}

func (r *CartRepo) CheckProduct(ctx context.Context, productID uuid.UUID) error {
	product := models.CartProduct{}
	err := r.db.QueryRow(ctx, getProduct, productID).Scan(&product.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrPoductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return err
	}

	return nil
}

func (r *CartRepo) AddProduct(ctx context.Context, cart models.Cart, product models.CartProduct) (models.Cart, error) {
	_, err := r.db.Exec(ctx, addProduct, cart.Id, product.Id, product.Quantity)
	if err != nil {
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return cart, err
	}

	cart, err = r.ReadCartProducts(ctx, cart)

	return cart, err
}

func (r *CartRepo) DeleteProduct(ctx context.Context, cart models.Cart, product models.CartProduct) (models.Cart, error) {
	_, err := r.db.Exec(ctx, deleteProduct, cart.Id, product.Id)
	if err != nil {
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return cart, err
	}

	cart, err = r.ReadCartProducts(ctx, cart)

	return cart, err
}

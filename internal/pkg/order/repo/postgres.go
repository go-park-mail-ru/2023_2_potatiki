package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	createOrder     = "INSERT INTO order_info (id, profile_id, delivery_at) VALUES ($1, $2, $3);"
	createOrderItem = "INSERT INTO order_item (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4);"
	getProductPrice = "SELECT price FROM product WHERE id=$1;"

	getCart     = "SELECT id FROM cart WHERE Profile_id=$1 and is_current = true;"
	getProducts = "SELECT p._id , p.name_product, p.Description, p.Price, p.ImgSrc, p.Rating, sc.Quantity " +
		"FROM shopping_cart_item sc JOIN product p ON sc.product_id=p.id WHERE p.id=$1;"
	getProduct = "SELECT id FROM product where id=$1;"
	addProduct = "insert into shopping_cart_item(cart_id, product_id, quantity) values ($1, $2, $3)" +
		" ON CONFLICT ON CONSTRAINT uq_shopping_cart_item_cart_id_product_id " +
		"do update set quantity=$3 WHERE shopping_cart_item.cart_id=$1 and shopping_cart_item.product_id=$2;"
	deleteProduct = "DELETE FROM shopping_cart_item WHERE cart_id=$1 and product_id=$2;"
)

var (
	ErrCartNotFound    = errors.New("cart not found")
	ErrPoductsNotFound = errors.New("products not found")
	ErrPoductNotFound  = errors.New("product not found")
)

type OrderRepo struct {
	db pgxtype.Querier
}

func NewOrderRepo(db pgxtype.Querier) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, cart models.Cart, userID uuid.UUID) (models.Order, error) {
	orderID := uuid.New()
	_, err := r.db.Exec(ctx, createOrder, orderID, userID, time.Now().Add(24*time.Hour))
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return models.Order{}, err
	}
	for _, cartProduct := range cart.Products {
		err = r.db.QueryRow(ctx, getProductPrice, cartProduct.Id).Scan(&cartProduct.Price)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return models.Order{}, ErrPoductNotFound
			}
			err = fmt.Errorf("error happened in row.Scan: %w", err)

			return models.Order{}, err
		}

		orderItemID := uuid.New()
		_, err = r.db.Exec(ctx, createOrderItem,
			orderItemID, orderID, cartProduct.Id, cartProduct.Quantity, cartProduct.Price)

		if err != nil {
			err = fmt.Errorf("error happened in db.Exec: %w", err)

			return models.Order{}, err
		}

	}

	return models.Order{}, nil // TODO: FIX
}

//func (r *CartRepo) CheckCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
//	cart := models.Cart{}
//	err := r.db.QueryRow(ctx, getCart, userID).Scan(&cart.Id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return models.Cart{}, ErrCartNotFound
//		}
//		err = fmt.Errorf("error happened in row.Scan: %w", err)
//
//		return models.Cart{}, err
//	}
//	return cart, nil
//}
//
//func (r *CartRepo) ReadCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
//	cart, err := r.CheckCart(ctx, userID)
//	if err != nil {
//		return models.Cart{}, err
//	}
//
//	cart, err = r.ReadCartProducts(ctx, cart)
//
//	return cart, err
//}
//
//func (r *CartRepo) ReadCartProducts(ctx context.Context, cart models.Cart) (models.Cart, error) {
//	rows, err := r.db.Query(ctx, getProducts, cart.Id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return cart, ErrPoductsNotFound
//		}
//		err = fmt.Errorf("error happened in db.QueryContext: %w", err)
//
//		return cart, err
//	}
//
//	product := models.CartProduct{}
//	for rows.Next() {
//		err = rows.Scan(
//			&product.Id,
//			&product.Name,
//			&product.Description,
//			&product.Price,
//			&product.ImgSrc,
//			&product.Rating,
//			&product.Quantity,
//		)
//		if err != nil {
//			err = fmt.Errorf("error happened in rows.Scan: %w", err)
//
//			return cart, err
//		}
//		cart.Products = append(cart.Products, product)
//	}
//	defer rows.Close()
//
//	return cart, nil
//}

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
	createAddress = `INSERT INTO address (id, profile_id, city, street, house, flat, is_current) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	updateAddress = `UPDATE address
	SET city = $1, street = $2, house = $3, flat = $4, is_current = $5 WHERE id = $6 AND profile_id = $7;`

	deleteAddress = `UPDATE address
	SET is_deleted = true WHERE id = $1 AND profile_id = $2 AND is_current = false;`

	makeCurrentAddress = `UPDATE address
	SET is_current = true WHERE id = $1 AND profile_id = $2;`

	getAddress = `SELECT city, street, house, flat, is_current
	FROM address
	WHERE id = $1 AND profile_id = $2;
	`

	createOrderItem = "INSERT INTO order_item (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4, $5);"

	getProductInfo = `
	SELECT p.name, p.description, p.price, p.imgsrc, p.rating, 
    c.id AS category_id, c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	WHERE p.id = $1;
	`

	getCurrentOrder = `
	SELECT p.id AS product_id, p.name AS product_name, p.description AS product_description, p.price AS product_price, 
	p.imgsrc AS product_imgsrc, p.rating AS product_rating, oi.quantity AS product_quantity, c.id AS category_id, 
	c.name AS category_name, o.status_id AS status_id
	FROM order_item oi
	JOIN product p ON oi.product_id = p.id
	JOIN order_info o ON oi.order_id = o.id
	JOIN category c ON p.category_id = c.id
	WHERE oi.order_id = $1;
	`

	getCurrentOrderID = "SELECT oi.id AS order_id " +
		"FROM order_info oi " +
		"WHERE oi.profile_id = $1 " +
		"ORDER BY oi.creation_at DESC;"

	getOrdersID = `
	SELECT id AS order_id
	FROM order_info
	WHERE profile_id = $1
	ORDER BY creation_at DESC;
	`
)

var (
	ErrAddressNotFound = errors.New("address not found")
)

type AddressRepo struct {
	db pgxtype.Querier
}

func NewAddressRepo(db pgxtype.Querier) *AddressRepo {
	return &AddressRepo{
		db: db,
	}
}

func (r *AddressRepo) CreateAddress(ctx context.Context, userID uuid.UUID, addressInfo models.AddressInfo) (models.Address, error) {
	addressID := uuid.NewV4()
	_, err := r.db.Exec(ctx, createAddress, addressID, userID,
		addressInfo.City,
		addressInfo.Street,
		addressInfo.House,
		addressInfo.Flat,
		true,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return models.Address{}, err
	}

	return models.Address{
		Id:        addressID,
		City:      addressInfo.City,
		Street:    addressInfo.Street,
		House:     addressInfo.House,
		Flat:      addressInfo.Flat,
		IsCurrent: true,
	}, nil
}

func (r *AddressRepo) UpdateAddress(ctx context.Context, addressInfo models.Address) error {
	_, err := r.db.Exec(ctx, updateAddress,
		addressInfo.City,
		addressInfo.Street,
		addressInfo.House,
		addressInfo.Flat,
		addressInfo.IsCurrent,
		addressInfo.Id,
		addressInfo.ProfileId,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *AddressRepo) DeleteAddress(ctx context.Context, addressInfo models.AddressDelete) error {
	_, err := r.db.Exec(ctx, deleteAddress,
		addressInfo.Id,
		addressInfo.ProfileId,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *AddressRepo) MakeCurrentAddress(ctx context.Context, addressInfo models.AddressMakeCurrent) error {
	_, err := r.db.Exec(ctx, makeCurrentAddress,
		addressInfo.Id,
		addressInfo.ProfileId,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *AddressRepo) GetAddress(ctx context.Context, userID uuid.UUID, addressID uuid.UUID) (models.Address, error) {
	address := models.Address{Id: addressID}
	err := r.db.QueryRow(ctx, getAddress, addressID, userID).
		Scan(
			&address.City,
			&address.Street,
			&address.House,
			&address.Flat,
			&address.IsCurrent,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Address{}, ErrAddressNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Address{}, err
	}

	return address, nil
}

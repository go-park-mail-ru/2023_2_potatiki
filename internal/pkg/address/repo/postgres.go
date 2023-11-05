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
	SET city = $1, street = $2, house = $3, flat = $4, is_current = $5 WHERE id = $6 AND profile_id = $7;` // FIX

	deleteAddress = `UPDATE address
	SET is_deleted = true WHERE id = $1 AND profile_id = $2 AND is_current = false;`

	makeCurrentAddress = `UPDATE address
	SET is_current = true WHERE id = $1 AND profile_id = $2;`

	readAddress = `SELECT city, street, house, flat, is_current
	FROM address
	WHERE id = $1 AND profile_id = $2;
	`

	readCurrentAddress = `SELECT id, city, street, house, flat, is_current
	FROM address
	WHERE profile_id = $1 AND is_current = true;
	`

	readAddresses = `SELECT id, city, street, house, flat, is_current
	FROM address
	WHERE profile_id = $1 AND is_deleted = false;
	`
)

var (
	ErrAddressNotFound        = errors.New("address not found")
	ErrAddressesNotFound      = errors.New("address not found")
	ErrCurrentAddressNotFound = errors.New("current address not found")
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
	result, err := r.db.Exec(ctx, makeCurrentAddress,
		addressInfo.Id,
		addressInfo.ProfileId,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}
	if result.RowsAffected() == 0 {
		return ErrCurrentAddressNotFound
	}

	return nil
}

func (r *AddressRepo) ReadAddress(ctx context.Context, userID uuid.UUID, addressID uuid.UUID) (models.Address, error) {
	address := models.Address{Id: addressID}
	err := r.db.QueryRow(ctx, readAddress, addressID, userID).
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

func (r *AddressRepo) ReadCurrentAddress(ctx context.Context, userID uuid.UUID) (models.Address, error) {
	address := models.Address{}
	err := r.db.QueryRow(ctx, readCurrentAddress, userID).
		Scan(
			&address.Id,
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

func (r *AddressRepo) ReadAllAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	rows, err := r.db.Query(ctx, readAddresses, userID)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Address{}, ErrAddressesNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Address{}, err
	}

	address := models.Address{}
	addresses := make([]models.Address, 0)
	for rows.Next() {
		err = rows.Scan(
			&address.Id,
			&address.City,
			&address.Street,
			&address.House,
			&address.Flat,
			&address.IsCurrent,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Address{}, err
		}
		addresses = append(addresses, address)
	}
	if len(addresses) == 0 {
		return []models.Address{}, ErrAddressesNotFound
	}

	return addresses, nil
}

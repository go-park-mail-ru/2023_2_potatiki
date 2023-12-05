package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
)

const (
	readPromocode = "SELECT * FROM promocode WHERE name=$1;"
)

type PromoRepo struct {
	db pgxtype.Querier // TODO: add logger
}

func NewPromoRepo(db pgxtype.Querier) *PromoRepo {
	return &PromoRepo{
		db: db,
	}
}

func (r *PromoRepo) ReadPromocode(ctx context.Context, promocodeName string) (*models.Promocode, error) {
	p := &models.Promocode{}
	if err := r.db.QueryRow(ctx, readPromocode, promocodeName).
		Scan(&p.Id, &p.Discount, &p.Name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &models.Promocode{}, promo.ErrPromocodeNotFound
		}
		return &models.Promocode{}, fmt.Errorf("error happened in row.Scan: %w", err)
	}
	return p, nil
}

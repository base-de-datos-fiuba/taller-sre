package products

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type PriceAudit struct {
	ID        int64     `json:"id"`
	OldPrice  string    `json:"old_price"`
	NewPrice  string    `json:"new_price"`
	ChangedAt time.Time `json:"changed_at"`
	ChangedBy string    `json:"changed_by"`
}

type PriceChange struct {
	ProductID int64       `json:"product_id"`
	Price     string      `json:"price"`
	Audit     *PriceAudit `json:"audit"`
}

func (r *Repository) ChangePrice(ctx context.Context, id int64, price string) (PriceChange, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return PriceChange{}, err
	}
	defer tx.Rollback(ctx)

	var oldPrice string
	if err := tx.QueryRow(ctx,
		"SELECT price::text FROM product WHERE product_id = $1 FOR UPDATE", id,
	).Scan(&oldPrice); err != nil {
		return PriceChange{}, err
	}

	change := PriceChange{ProductID: id, Price: oldPrice}
	var changed bool
	if err := tx.QueryRow(ctx,
		"SELECT $1::numeric IS DISTINCT FROM $2::numeric", oldPrice, price,
	).Scan(&changed); err != nil {
		return PriceChange{}, err
	}
	if changed {
		if err := tx.QueryRow(ctx,
			"UPDATE product SET price = $2::numeric WHERE product_id = $1 RETURNING price::text",
			id, price,
		).Scan(&change.Price); err != nil {
			return PriceChange{}, err
		}
		change.Audit = &PriceAudit{}
		if err := tx.QueryRow(ctx, `
			SELECT audit_id, old_price::text, new_price::text, changed_at, changed_by
			FROM product_price_audit
			WHERE product_id = $1
			ORDER BY audit_id DESC
			LIMIT 1`, id,
		).Scan(&change.Audit.ID, &change.Audit.OldPrice, &change.Audit.NewPrice,
			&change.Audit.ChangedAt, &change.Audit.ChangedBy); err != nil {
			return PriceChange{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return PriceChange{}, err
	}
	return change, nil
}

func IsProductNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

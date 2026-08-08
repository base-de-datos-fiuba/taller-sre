package products

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List(ctx context.Context) ([]Product, error) {
	rows, err := r.db.Query(ctx, `
		SELECT product_id, name, COALESCE(description, ''), price::text, stock
		FROM product
		ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProducts(rows)
}

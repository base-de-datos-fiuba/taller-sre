package products

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// SearchVulnerable is INTENTIONALLY VULNERABLE.
// This code exists only for the university SQL injection demonstration.
// Never write production code like this.
func (r *Repository) SearchVulnerable(ctx context.Context, search string) ([]Product, error) {
	// Deliberate SQL injection: untrusted input is concatenated into SQL.
	query := `SELECT product_id, name, COALESCE(description, ''), price::text, stock
		FROM product WHERE name ILIKE '%` + search + `%' ORDER BY name`

	// Simple protocol is deliberate: it permits stacked statements for the lesson.
	rows, err := r.db.Query(ctx, query, pgx.QueryExecModeSimpleProtocol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanProducts(rows)
}

type productRows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
}

func scanProducts(rows productRows) ([]Product, error) {
	products := make([]Product, 0)
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

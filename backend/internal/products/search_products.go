package products

import "context"

// Search treats the external value as data by sending it separately from the
// SQL statement. PostgreSQL never parses the search text as SQL instructions.
func (r *Repository) Search(ctx context.Context, search string) ([]Product, error) {
	query := `SELECT product_id, name, COALESCE(description, ''), price::text, stock
		FROM product
		WHERE name ILIKE $1
		ORDER BY name`

	rows, err := r.db.Query(ctx, query, "%"+search+"%")
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

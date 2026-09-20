package products

import "context"

func (r *Repository) ChangePrice(ctx context.Context, id int64, price string) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE product SET price = $2::numeric WHERE product_id = $1", id, price,
	)
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}

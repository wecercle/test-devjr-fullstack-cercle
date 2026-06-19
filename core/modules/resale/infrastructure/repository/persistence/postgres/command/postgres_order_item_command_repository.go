package command

import (
    "context"
    "database/sql"

    cmdRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
)

type PostgresOrderItemCommandRepository struct {
    db *sql.DB
}

func NewPostgresOrderItemCommandRepository(db *sql.DB) *PostgresOrderItemCommandRepository {
    return &PostgresOrderItemCommandRepository{db: db}
}

// UpdateShippingStatus implements cmdRepo.OrderItemCommandRepository
func (r *PostgresOrderItemCommandRepository) UpdateShippingStatus(ctx context.Context, orderID string, itemID string, status string) error {
    const stmt = `
UPDATE cercle_test.resale_order_item
SET shipping_status = $3,
    updated_at = NOW()
WHERE fk_resale_order_id = $1
  AND id = $2;
`
    res, err := r.db.ExecContext(ctx, stmt, orderID, itemID, status)
    if err != nil {
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return sql.ErrNoRows
    }
    return nil
}

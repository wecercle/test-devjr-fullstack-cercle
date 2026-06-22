package query

import (
    "context"
    "database/sql"
    "errors"

    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/model"
    queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type PostgresOrderItemQueryRepository struct {
    db *sql.DB
}

func NewPostgresOrderItemQueryRepository(db *sql.DB) *PostgresOrderItemQueryRepository {
    return &PostgresOrderItemQueryRepository{db: db}
}

// ListByCPFAndOrderID implements queryRepo.OrderItemQueryRepository
func (r *PostgresOrderItemQueryRepository) ListByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*model.OrderItem, error) {
    const stmt = `
SELECT
    roi.id,
    roi.fk_resale_order_id,
    roi.sku,
    roi.name,
    roi.quantity,
    roi.amount_value,
    roi.shipping_code,
    roi.shipping_status,
    roi.delivered_at,
    roi.deleted_at
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE ro.id = $1 AND u.document_number = $2
ORDER BY roi.created_at ASC;
`
    rows, err := r.db.QueryContext(ctx, stmt, orderID, cpf)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []*model.OrderItem
    for rows.Next() {
        var it model.OrderItem
        if err := rows.Scan(
            &it.ID,
            &it.FKResaleOrderID,
            &it.SKU,
            &it.Name,
            &it.Quantity,
            &it.AmountValue,
            &it.ShippingCode,
            &it.ShippingStatus,
            &it.DeliveredAt,
            &it.DeletedAt,
        ); err != nil {
            return nil, err
        }
        items = append(items, &it)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    if len(items) == 0 {
        return nil, sql.ErrNoRows
    }
    return items, nil
}

// GetByOrderAndItemID implements queryRepo.OrderItemQueryRepository
func (r *PostgresOrderItemQueryRepository) GetByOrderAndItemID(ctx context.Context, orderID, itemID string) (*model.OrderItem, error) {
    const stmt = `
SELECT
    roi.id,
    roi.fk_resale_order_id,
    roi.sku,
    roi.name,
    roi.quantity,
    roi.amount_value,
    roi.shipping_code,
    roi.shipping_status,
    roi.delivered_at,
    roi.deleted_at
FROM cercle_test.resale_order_item roi
WHERE roi.fk_resale_order_id = $1 AND roi.id = $2;
`
    row := r.db.QueryRowContext(ctx, stmt, orderID, itemID)
    var it model.OrderItem
    if err := row.Scan(
        &it.ID,
        &it.FKResaleOrderID,
        &it.SKU,
        &it.Name,
        &it.Quantity,
        &it.AmountValue,
        &it.ShippingCode,
        &it.ShippingStatus,
        &it.DeliveredAt,
        &it.DeletedAt,
    ); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, sql.ErrNoRows
        }
        return nil, err
    }
    if it.DeletedAt != nil {
        return nil, sql.ErrNoRows
    }
    return &it, nil
}

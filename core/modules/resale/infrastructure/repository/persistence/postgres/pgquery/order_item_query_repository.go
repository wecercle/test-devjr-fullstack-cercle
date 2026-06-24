package pgquery

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type OrderItemQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewOrderItemQueryRepository(querier *databaseQuery.Queries) *OrderItemQueryRepository {
	return &OrderItemQueryRepository{querier: querier}
}

func (r *OrderItemQueryRepository) SelectItemsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*aggregate.OrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}

	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(ctx, databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	items := make([]*aggregate.OrderItem, 0, len(rows))
	for _, row := range rows {
		item, err := aggregate.ReconstituteOrderItem(
			row.ID.String(),
			row.FkResaleOrderID.String(),
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			row.ShippingCode,
			row.ShippingStatus,
			row.DeliveredAt,
			row.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *OrderItemQueryRepository) SelectItemByID(ctx context.Context, cpf string, orderID string, itemID string) (*aggregate.OrderItem, error) {
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, domainException.ErrInvalidItemID
	}
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}

	row, err := r.querier.SelectOrderItemByID(ctx, databaseQuery.SelectOrderItemByIDParams{
		ID:             parsedItemID,
		ResaleOrderID:  parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainException.ErrOrderItemNotFound
		}
		return nil, err
	}

	item, err := aggregate.ReconstituteOrderItem(
		row.ID.String(),
		row.FkResaleOrderID.String(),
		row.Sku,
		row.Name,
		row.Quantity,
		row.AmountValue,
		row.ShippingCode,
		row.ShippingStatus,
		row.DeliveredAt,
		row.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return item, nil
}

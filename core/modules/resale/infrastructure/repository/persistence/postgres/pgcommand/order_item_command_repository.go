package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type OrderItemCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewOrderItemCommandRepository(querier *databaseQuery.Queries) *OrderItemCommandRepository {
	return &OrderItemCommandRepository{querier: querier}
}

func (r *OrderItemCommandRepository) CancelItem(ctx context.Context, resaleOrderID string, itemID string) error {
	parsedOrderID, err := uuid.Parse(resaleOrderID)
	if err != nil {
		return domainException.ErrInvalidOrderID
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return domainException.ErrInvalidItemID
	}

	rowsAffected, err := r.querier.UpdateOrderItemShippingStatus(ctx, databaseQuery.UpdateOrderItemShippingStatusParams{
		ShippingStatus: sql.NullString{String: "RETURNED", Valid: true},
		ResaleOrderID:  parsedOrderID,
		ID:             parsedItemID,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domainException.ErrOrderItemNotFound
	}

	return nil
}

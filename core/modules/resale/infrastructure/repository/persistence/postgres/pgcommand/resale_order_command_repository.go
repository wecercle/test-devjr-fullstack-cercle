package pgcommand

import (
	"context"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
)

type ResaleOrderCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleOrderCommandRepository(querier *databaseQuery.Queries) *ResaleOrderCommandRepository {
	return &ResaleOrderCommandRepository{querier: querier}
}

func (r *ResaleOrderCommandRepository) MarkOrderItemAsReturned(ctx context.Context, orderID, itemID string) (bool, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return false, err
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := r.querier.MarkOrderItemAsReturned(ctx, databaseQuery.MarkOrderItemAsReturnedParams{
		ResaleOrderID: parsedOrderID,
		ID:            parsedItemID,
	})
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

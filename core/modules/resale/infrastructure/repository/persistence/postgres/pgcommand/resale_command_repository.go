package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
)

// ResaleCommandRepository é a implementação concreta (Postgres) da interface
// command.ResaleCommandRepository.
type ResaleCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleCommandRepository(querier *databaseQuery.Queries) *ResaleCommandRepository {
	return &ResaleCommandRepository{querier: querier}
}

func (r *ResaleCommandRepository) UpdateShippingStatus(ctx context.Context, orderID, itemID, status string) (int64, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return 0, err
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return 0, err
	}

	return r.querier.UpdateOrderItemShippingStatus(ctx, databaseQuery.UpdateOrderItemShippingStatusParams{
		ShippingStatus: sql.NullString{String: status, Valid: true},
		ResaleOrderID:  parsedOrderID,
		ID:             parsedItemID,
	})
}

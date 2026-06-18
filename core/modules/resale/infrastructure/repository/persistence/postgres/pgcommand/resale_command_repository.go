package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type ResaleCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleCommandRepository(querier *databaseQuery.Queries) *ResaleCommandRepository {
	return &ResaleCommandRepository{querier: querier}
}

func (r *ResaleCommandRepository) UpdateOrderItemShippingStatus(ctx context.Context, item *aggregate.OrderItem) error {
	parsedOrderID, err := uuid.Parse(item.ResaleOrderID())
	if err != nil {
		return err
	}
	parsedItemID, err := uuid.Parse(item.ID())
	if err != nil {
		return err
	}

	rowsAffected, err := r.querier.UpdateOrderItemShippingStatus(ctx, databaseQuery.UpdateOrderItemShippingStatusParams{
		ShippingStatus: sql.NullString{String: item.ShippingStatus(), Valid: true},
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

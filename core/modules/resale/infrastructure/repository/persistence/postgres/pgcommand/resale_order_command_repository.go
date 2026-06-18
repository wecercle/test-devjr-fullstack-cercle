package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type ResaleOrderCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleOrderCommandRepository(querier *databaseQuery.Queries) *ResaleOrderCommandRepository {
	return &ResaleOrderCommandRepository{querier: querier}
}

func (r *ResaleOrderCommandRepository) UpdateItemShippingStatus(ctx context.Context, item *aggregate.ResaleOrderItem) error {
	parsedOrderID, err := uuid.Parse(item.FkResaleOrderID())
	if err != nil {
		return domainException.ErrInvalidOrderID
	}
	parsedItemID, err := uuid.Parse(item.ID())
	if err != nil {
		return domainException.ErrInvalidOrderItemID
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

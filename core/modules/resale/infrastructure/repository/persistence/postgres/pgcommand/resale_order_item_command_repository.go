package pgcommand

import (
	"context"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type ResaleOrderItemCommandRepository struct{ querier *databaseQuery.Queries }

func NewResaleOrderItemCommandRepository(q *databaseQuery.Queries) *ResaleOrderItemCommandRepository {
	return &ResaleOrderItemCommandRepository{querier: q}
}

func (r *ResaleOrderItemCommandRepository) UpdateShippingStatus(ctx context.Context, id, resaleOrderID, shippingStatus string) error {
	parsedID, _ := uuid.Parse(id)
	parsedOrderID, _ := uuid.Parse(resaleOrderID)
	rows, err := r.querier.UpdateOrderItemShippingStatus(ctx, databaseQuery.UpdateOrderItemShippingStatusParams{
		ShippingStatus: shippingStatus, ResaleOrderID: parsedOrderID, ID: parsedID,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return domainException.ErrItemNotFound
	}
	return nil
}

package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type OrderCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewOrderCommandRepository(querier *databaseQuery.Queries) *OrderCommandRepository {
	return &OrderCommandRepository{querier: querier}
}

func (r *OrderCommandRepository) UpdateOrderItemShippingStatus(ctx context.Context, resaleOrderID string, id string, shippingStatus string) error {
	parsedResaleOrderID, err := uuid.Parse(resaleOrderID)
	if err != nil {
		return err
	}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	rowsAffected, err := r.querier.UpdateOrderItemShippingStatus(ctx, databaseQuery.UpdateOrderItemShippingStatusParams{
		ShippingStatus: sql.NullString{String: shippingStatus, Valid: true},
		ResaleOrderID:  parsedResaleOrderID,
		ID:             parsedID,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domainException.ErrOrderItemNotFound
	}
	return nil
}

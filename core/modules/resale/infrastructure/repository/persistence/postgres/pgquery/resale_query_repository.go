package pgquery

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository"
)

type ResaleQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleQueryRepository(querier *databaseQuery.Queries) repository.ResaleQueryRepository {
	return &ResaleQueryRepository{querier: querier}
}

func (r *ResaleQueryRepository) GetOrderItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error) {
	uid, _ := uuid.Parse(orderID)

	arg := databaseQuery.GetOrderItemsByCPFAndOrderIDParams{
		DocumentNumber: sql.NullString{String: cpf, Valid: cpf != ""}, 
		OrderID:        uid,
	}

	rows, err := r.querier.GetOrderItemsByCPFAndOrderID(ctx, arg)
	if err != nil {
		return nil, err
	}

	var items []*aggregate.ResaleOrderItem
	for _, row := range rows {
		items = append(items, &aggregate.ResaleOrderItem{
			ID:              row.ID.String(), 
			FkResaleOrderID: row.FkResaleOrderID.String(),
			Sku:             row.Sku,
			Name:            row.Name,
			Quantity:        int(row.Quantity),
		})
	}
	return items, nil
}

func (r *ResaleQueryRepository) GetOrderItemForCancellation(ctx context.Context, cpf, orderID, itemID string) (*aggregate.ResaleOrderItem, error) {
	orderUID, _ := uuid.Parse(orderID)
	itemUID, _ := uuid.Parse(itemID)

	arg := databaseQuery.GetOrderItemForCancellationParams{
		DocumentNumber: sql.NullString{String: cpf, Valid: cpf != ""},
		OrderID:        orderUID,
		ItemID:         itemUID,
	}

	row, err := r.querier.GetOrderItemForCancellation(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &aggregate.ResaleOrderItem{
		ID: row.ID.String(),
	}, nil
}
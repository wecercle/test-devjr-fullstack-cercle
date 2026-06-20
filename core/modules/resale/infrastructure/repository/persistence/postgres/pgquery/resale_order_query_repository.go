package pgquery

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/valueobject"
	sharedInfrastructure "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/infrastructure"
)

type ResaleOrderQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleOrderQueryRepository(querier *databaseQuery.Queries) *ResaleOrderQueryRepository {
	return &ResaleOrderQueryRepository{querier: querier}
}

func (r *ResaleOrderQueryRepository) ExistsOrderByCPFAndOrderID(ctx context.Context, cpf, orderID string) (bool, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return false, err
	}

	return r.querier.ExistsOrderByCPFAndOrderID(ctx, databaseQuery.ExistsOrderByCPFAndOrderIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
}

func (r *ResaleOrderQueryRepository) SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}

	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(ctx, databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	items := make([]*aggregate.ResaleOrderItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, aggregate.NewResaleOrderItem(
			row.ID.String(),
			row.FkResaleOrderID.String(),
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			nullStringToPointer(row.ShippingCode),
			valueobject.ShippingStatus(row.ShippingStatus.String),
			nil,
		))
	}

	return items, nil
}

func (r *ResaleOrderQueryRepository) SelectOrderItemByOrderIDAndItemID(ctx context.Context, orderID, itemID string) (*aggregate.ResaleOrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, err
	}

	row, err := r.querier.SelectOrderItemByOrderIDAndItemID(ctx, databaseQuery.SelectOrderItemByOrderIDAndItemIDParams{
		OrderID: parsedOrderID,
		ItemID:  parsedItemID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainException.ErrOrderItemNotFound
		}
		return nil, err
	}

	return aggregate.NewResaleOrderItem(
		row.ID.String(),
		row.FkResaleOrderID.String(),
		row.Sku,
		row.Name,
		row.Quantity,
		row.AmountValue,
		nullStringToPointer(row.ShippingCode),
		valueobject.ShippingStatus(row.ShippingStatus.String),
		sharedInfrastructure.NullTimeToPointer(row.DeliveredAt),
	), nil
}

func nullStringToPointer(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

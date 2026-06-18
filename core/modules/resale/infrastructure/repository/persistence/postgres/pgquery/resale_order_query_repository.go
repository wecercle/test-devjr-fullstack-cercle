package pgquery

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedInfrastructure "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/infrastructure"
)

type ResaleOrderQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleOrderQueryRepository(querier *databaseQuery.Queries) *ResaleOrderQueryRepository {
	return &ResaleOrderQueryRepository{querier: querier}
}

func (r *ResaleOrderQueryRepository) SelectOrderByCPFAndOrderID(ctx context.Context, cpf, orderID string) error {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return domainException.ErrInvalidOrderID
	}

	_, err = r.querier.SelectResaleOrderByCPFAndOrderID(ctx, databaseQuery.SelectResaleOrderByCPFAndOrderIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return domainException.ErrOrderNotFound
		}
		return err
	}

	return nil
}

func (r *ResaleOrderQueryRepository) SelectItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error) {
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

	items := make([]*aggregate.ResaleOrderItem, 0, len(rows))
	for _, row := range rows {
		item, err := aggregate.ReconstituteResaleOrderItem(
			row.ID.String(),
			row.FkResaleOrderID.String(),
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			nullableString(row.ShippingCode),
			nullableString(row.ShippingStatus),
			sharedInfrastructure.NullTimeToPointer(row.DeliveredAt),
			row.CreatedAt,
			row.UpdatedAt,
			sharedInfrastructure.NullTimeToPointer(row.DeletedAt),
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ResaleOrderQueryRepository) SelectItemByCPFOrderIDAndItemID(ctx context.Context, cpf, orderID, itemID string) (*aggregate.ResaleOrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderItemID
	}

	row, err := r.querier.SelectOrderItemByCPFOrderIDAndItemID(ctx, databaseQuery.SelectOrderItemByCPFOrderIDAndItemIDParams{
		OrderID:        parsedOrderID,
		ItemID:         parsedItemID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainException.ErrOrderItemNotFound
		}
		return nil, err
	}

	return aggregate.ReconstituteResaleOrderItem(
		row.ID.String(),
		row.FkResaleOrderID.String(),
		row.Sku,
		row.Name,
		row.Quantity,
		row.AmountValue,
		nullableString(row.ShippingCode),
		nullableString(row.ShippingStatus),
		sharedInfrastructure.NullTimeToPointer(row.DeliveredAt),
		row.CreatedAt,
		row.UpdatedAt,
		sharedInfrastructure.NullTimeToPointer(row.DeletedAt),
	)
}

func nullableString(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.String
}

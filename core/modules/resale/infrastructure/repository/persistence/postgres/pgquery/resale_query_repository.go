package pgquery

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedInfrastructure "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/infrastructure"
)

type ResaleQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleQueryRepository(querier *databaseQuery.Queries) *ResaleQueryRepository {
	return &ResaleQueryRepository{querier: querier}
}

func (r *ResaleQueryRepository) OrderExistsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) (bool, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return false, err
	}

	_, err = r.querier.SelectOrderByCPFAndOrderID(ctx, databaseQuery.SelectOrderByCPFAndOrderIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *ResaleQueryRepository) SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*aggregate.OrderItem, error) {
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

	items := make([]*aggregate.OrderItem, 0, len(rows))
	for _, row := range rows {
		item, err := buildOrderItem(
			row.ID,
			row.FkResaleOrderID,
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			row.ShippingCode,
			row.ShippingStatus,
			row.DeliveredAt,
			row.CreatedAt,
			row.UpdatedAt,
			row.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ResaleQueryRepository) SelectOrderItemByCPFOrderIDAndItemID(ctx context.Context, cpf string, orderID string, itemID string) (*aggregate.OrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, err
	}

	row, err := r.querier.SelectOrderItemByCPFOrderIDAndItemID(ctx, databaseQuery.SelectOrderItemByCPFOrderIDAndItemIDParams{
		OrderID:        parsedOrderID,
		DocumentNumber: sql.NullString{String: cpf, Valid: true},
		ItemID:         parsedItemID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainException.ErrOrderItemNotFound
		}
		return nil, err
	}

	return buildOrderItem(
		row.ID,
		row.FkResaleOrderID,
		row.Sku,
		row.Name,
		row.Quantity,
		row.AmountValue,
		row.ShippingCode,
		row.ShippingStatus,
		row.DeliveredAt,
		row.CreatedAt,
		row.UpdatedAt,
		row.DeletedAt,
	)
}

func buildOrderItem(
	id uuid.UUID,
	resaleOrderID uuid.UUID,
	sku string,
	name string,
	quantity int32,
	amountValue string,
	shippingCode sql.NullString,
	shippingStatus sql.NullString,
	deliveredAt sql.NullTime,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt sql.NullTime,
) (*aggregate.OrderItem, error) {
	shippingCodeValue := ""
	if shippingCode.Valid {
		shippingCodeValue = shippingCode.String
	}

	shippingStatusValue := ""
	if shippingStatus.Valid {
		shippingStatusValue = shippingStatus.String
	}

	return aggregate.ReconstituteOrderItem(
		id.String(),
		resaleOrderID.String(),
		sku,
		name,
		quantity,
		amountValue,
		shippingCodeValue,
		shippingStatusValue,
		sharedInfrastructure.NullTimeToPointer(deliveredAt),
		createdAt,
		updatedAt,
		sharedInfrastructure.NullTimeToPointer(deletedAt),
	)
}

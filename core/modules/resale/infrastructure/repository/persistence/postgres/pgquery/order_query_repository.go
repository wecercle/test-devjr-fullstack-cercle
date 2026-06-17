package pgquery

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type OrderQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewOrderQueryRepository(querier *databaseQuery.Queries) *OrderQueryRepository {
	return &OrderQueryRepository{
		querier: querier,
	}
}

func (r *OrderQueryRepository) SelectOrderItemsByCPFAndOrderID(
	ctx context.Context,
	cpf string,
	orderID string,
) ([]*aggregate.OrderItem, error) {

	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}

	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(
		ctx,
		databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
			OrderID: parsedOrderID,
			DocumentNumber: sql.NullString{
				String: cpf,
				Valid:  true,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	items := make([]*aggregate.OrderItem, 0, len(rows))

	for _, row := range rows {

		shippingCode := ""
		if row.ShippingCode.Valid {
			shippingCode = row.ShippingCode.String
		}

		shippingStatus := ""
		if row.ShippingStatus.Valid {
			shippingStatus = row.ShippingStatus.String
		}

		var deliveredAtPtr *time.Time
		if row.DeliveredAt.Valid {
			d := row.DeliveredAt.Time
			deliveredAtPtr = &d
		}

		item, err := aggregate.ReconstituteOrderItem(
			row.ID.String(),
			row.FkResaleOrderID.String(),
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			shippingCode,
			shippingStatus,
			deliveredAtPtr,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

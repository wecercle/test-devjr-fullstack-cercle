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

type ResaleOrderItemQueryRepository struct{ querier *databaseQuery.Queries }

func NewResaleOrderItemQueryRepository(q *databaseQuery.Queries) *ResaleOrderItemQueryRepository {
	return &ResaleOrderItemQueryRepository{querier: q}
}

func (r *ResaleOrderItemQueryRepository) SelectByOrderIDAndCPF(ctx context.Context, orderID, cpf string) ([]*aggregate.ResaleOrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}
	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(ctx, databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
		OrderID: parsedOrderID, DocumentNumber: cpf,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*aggregate.ResaleOrderItem, 0, len(rows))
	for _, row := range rows {
		item, err := aggregate.ReconstituteResaleOrderItem(
			row.ID.String(), row.FkResaleOrderID.String(),
			row.Sku, row.Name, row.Quantity, row.AmountValue,
			row.ShippingCode, row.ShippingStatus, nil,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ResaleOrderItemQueryRepository) SelectByIDAndOrderIDAndCPF(ctx context.Context, id, resaleOrderID, cpf string) (*aggregate.ResaleOrderItem, error) {
	parsedID, _ := uuid.Parse(id)
	parsedOrderID, _ := uuid.Parse(resaleOrderID)
	row, err := r.querier.SelectOrderItemByIDAndOrderIDAndCPF(ctx, databaseQuery.SelectOrderItemByIDAndOrderIDAndCPFParams{
		ID: parsedID, ResaleOrderID: parsedOrderID, DocumentNumber: cpf,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainException.ErrItemNotFound
		}
		return nil, err
	}
	return aggregate.ReconstituteResaleOrderItem(
		row.ID.String(), row.FkResaleOrderID.String(),
		row.Sku, row.Name, row.Quantity, row.AmountValue,
		row.ShippingCode, row.ShippingStatus,
		sharedInfrastructure.NullTimeToPointer(row.DeliveredAt),
	)
}

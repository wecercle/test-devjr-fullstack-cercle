package pgquery

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/aggregate"
	sharedInfrastructure "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/infrastructure"
)

type ResaleQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleQueryRepository(querier *databaseQuery.Queries) *ResaleQueryRepository {
	return &ResaleQueryRepository{querier: querier}
}

func (r *ResaleQueryRepository) SelectItemsByOrderAndCPF(ctx context.Context, orderID, cpf string) ([]*aggregate.ResaleOrderItem, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		// UUID malformado nunca vai encontrar nada no banco; tratamos como
		// "lista vazia", e o use case já transforma lista vazia em 404.
		return []*aggregate.ResaleOrderItem{}, nil
	}

	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(ctx, databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
		OrderID: parsedOrderID,
		DocumentNumber: sql.NullString{
			String: cpf,
			Valid:  cpf != "",
		},
	})
	if err != nil {
		return nil, err
	}

	items := make([]*aggregate.ResaleOrderItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, aggregate.ReconstituteResaleOrderItem(
			row.ID.String(),
			row.FkResaleOrderID.String(),
			row.Sku,
			row.Name,
			row.Quantity,
			row.AmountValue,
			row.ShippingCode.String,
			row.ShippingStatus.String,
			sharedInfrastructure.NullTimeToPointer(row.DeliveredAt),
			sharedInfrastructure.NullTimeToPointer(row.DeletedAt),
		))
	}
	return items, nil
}

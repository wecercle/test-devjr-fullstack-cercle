package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleOrderQueryRepository interface {
	SelectOrderByCPFAndOrderID(ctx context.Context, cpf, orderID string) error
	SelectItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error)
	SelectItemByCPFOrderIDAndItemID(ctx context.Context, cpf, orderID, itemID string) (*aggregate.ResaleOrderItem, error)
}

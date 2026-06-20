package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

// ResaleOrderQueryRepository defines read operations for resale orders and items.
type ResaleOrderQueryRepository interface {
	ExistsOrderByCPFAndOrderID(ctx context.Context, cpf, orderID string) (bool, error)
	SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error)
	SelectOrderItemByOrderIDAndItemID(ctx context.Context, orderID, itemID string) (*aggregate.ResaleOrderItem, error)
}

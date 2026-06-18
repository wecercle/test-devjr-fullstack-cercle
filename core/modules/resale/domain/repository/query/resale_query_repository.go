package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleQueryRepository interface {
	OrderExistsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) (bool, error)
	SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*aggregate.OrderItem, error)
	SelectOrderItemByCPFOrderIDAndItemID(ctx context.Context, cpf string, orderID string, itemID string) (*aggregate.OrderItem, error)
}

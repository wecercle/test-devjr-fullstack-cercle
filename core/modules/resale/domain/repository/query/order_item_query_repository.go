package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type OrderItemQueryRepository interface {
	SelectItemsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*aggregate.OrderItem, error)
	SelectItemByID(ctx context.Context, cpf string, orderID string, itemID string) (*aggregate.OrderItem, error)
}

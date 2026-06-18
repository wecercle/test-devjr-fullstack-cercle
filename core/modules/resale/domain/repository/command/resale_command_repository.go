package command

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleCommandRepository interface {
	UpdateOrderItemShippingStatus(ctx context.Context, item *aggregate.OrderItem) error
}

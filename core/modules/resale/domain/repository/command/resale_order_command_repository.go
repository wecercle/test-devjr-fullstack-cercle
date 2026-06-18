package command

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleOrderCommandRepository interface {
	UpdateItemShippingStatus(ctx context.Context, item *aggregate.ResaleOrderItem) error
}

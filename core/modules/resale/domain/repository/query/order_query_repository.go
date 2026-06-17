package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type OrderQueryRepository interface {
	SelectOrderItemsByCPFAndOrderID(
		ctx context.Context,
		cpf string,
		orderID string,
	) ([]*aggregate.OrderItem, error)
}

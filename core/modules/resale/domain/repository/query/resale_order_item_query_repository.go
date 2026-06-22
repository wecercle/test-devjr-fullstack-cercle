package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleOrderItemQueryRepository interface {
	SelectByOrderIDAndCPF(ctx context.Context, orderID, cpf string) ([]*aggregate.ResaleOrderItem, error)
	SelectByIDAndOrderIDAndCPF(ctx context.Context, id, resaleOrderID, cpf string) (*aggregate.ResaleOrderItem, error)
}

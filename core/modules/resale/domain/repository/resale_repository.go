package repository

import (
	"context"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)


type ResaleQueryRepository interface {
	GetOrderItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error)
	GetOrderItemForCancellation(ctx context.Context, cpf, orderID, itemID string) (*aggregate.ResaleOrderItem, error)
}


type ResaleCommandRepository interface {
	CancelOrderItem(ctx context.Context, itemID string) error
}
package command

import "context"

type OrderCommandRepository interface {
	UpdateOrderItemShippingStatus(ctx context.Context, resaleOrderID string, id string, shippingStatus string) error
}

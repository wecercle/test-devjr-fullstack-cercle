package command

import "context"

type ResaleOrderItemCommandRepository interface {
	UpdateShippingStatus(ctx context.Context, id, resaleOrderID, shippingStatus string) error
}

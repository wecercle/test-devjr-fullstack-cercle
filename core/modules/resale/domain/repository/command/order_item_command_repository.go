package command

import "context"

type OrderItemCommandRepository interface {
    // Update shipping status of an item belonging to a specific order
    UpdateShippingStatus(ctx context.Context, orderID string, itemID string, status string) error
}

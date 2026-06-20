package command

import "context"

// ResaleOrderCommandRepository defines write operations for resale order items.
type ResaleOrderCommandRepository interface {
	MarkOrderItemAsReturned(ctx context.Context, orderID, itemID string) (bool, error)
}

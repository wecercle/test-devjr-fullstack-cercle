package command

import (
	"context"
)

type OrderItemCommandRepository interface {
	CancelItem(ctx context.Context, resaleOrderID string, itemID string) error
}

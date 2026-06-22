package query

import "context"

type OrderItemQueryRepository interface {
    // Returns all items for given cpf and orderID
    ListByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]*OrderItem, error)
    // Retrieves single item by orderID and itemID
    GetByOrderAndItemID(ctx context.Context, orderID string, itemID string) (*OrderItem, error)
}

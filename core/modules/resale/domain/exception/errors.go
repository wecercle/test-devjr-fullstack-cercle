package exception

var (
	ErrInvalidOrderID         = &DomainError{Code: "invalid_order_id", Message: "invalid order id"}
	ErrInvalidOrderItemID     = &DomainError{Code: "invalid_order_item_id", Message: "invalid order item id"}
	ErrOrderNotFound          = &DomainError{Code: "order_not_found", Message: "order not found"}
	ErrOrderItemNotFound      = &DomainError{Code: "order_item_not_found", Message: "order item not found"}
	ErrOrderItemNotReturnable = &DomainError{Code: "order_item_not_returnable", Message: "order item is not returnable"}
	ErrOrderItemReturnExpired = &DomainError{Code: "order_item_return_expired", Message: "order item return period expired"}
)

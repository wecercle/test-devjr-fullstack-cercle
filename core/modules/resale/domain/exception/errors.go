package exception

var (
	ErrInvalidOrderID     = &DomainError{Code: "invalid_order_id", Message: "invalid order id"}
	ErrInvalidItemID      = &DomainError{Code: "invalid_item_id", Message: "invalid item id"}
	ErrInvalidCPF         = &DomainError{Code: "invalid_cpf", Message: "invalid cpf"}
	ErrOrderNotFound      = &DomainError{Code: "order_not_found", Message: "order not found or does not belong to the user"}
	ErrOrderItemNotFound  = &DomainError{Code: "order_item_not_found", Message: "order item not found or does not belong to the order"}
	ErrItemAlreadyReturned = &DomainError{Code: "item_already_returned", Message: "item is already returned"}
	ErrReturnWindowExpired = &DomainError{Code: "return_window_expired", Message: "return window of 7 days after delivery has expired"}
)

package exception

// Erros de domínio do módulo Resale
var (
	ErrOrderNotFound                = &DomainError{Code: "order_not_found", Message: "order not found"}
	ErrOrderItemNotFound            = &DomainError{Code: "order_item_not_found", Message: "order item not found"}
	ErrOrderItemAlreadyReturned     = &DomainError{Code: "order_item_already_returned", Message: "order item already returned"}
	ErrOrderItemCancelWindowExpired = &DomainError{Code: "order_item_cancel_window_expired", Message: "order item cancel window expired"}
)

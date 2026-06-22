package exception

// Erros de domínio do módulo Resale.
var (
	ErrInvalidOrderID           = &DomainError{Code: "invalid_order_id", Message: "invalid order id"}
	ErrInvalidItemID            = &DomainError{Code: "invalid_item_id", Message: "invalid item id"}
	ErrOrderNotFound            = &DomainError{Code: "order_not_found", Message: "order not found"}
	ErrItemNotFound             = &DomainError{Code: "item_not_found", Message: "item not found"}
	ErrItemNotEligibleForReturn = &DomainError{Code: "item_not_eligible_for_return", Message: "item is not eligible for return, delivery exceeded 7 days"}
)

package exception

var (
	ErrInvalidOrderID              = &DomainError{Code: "invalid_order_id", Message: "invalid order id"}
	ErrInvalidOrderItemID          = &DomainError{Code: "invalid_order_item_id", Message: "invalid order item id"}
	ErrInvalidUserDocumentNumber   = &DomainError{Code: "invalid_user_document_number", Message: "invalid user document number"}
	ErrOrderNotFound               = &DomainError{Code: "order_not_found", Message: "order not found"}
	ErrOrderItemNotFound           = &DomainError{Code: "order_item_not_found", Message: "order item not found"}
	ErrItemNotEligibleForReturn    = &DomainError{Code: "item_not_eligible_for_return", Message: "item is not eligible for return"}
	ErrCancelWindowExpired         = &DomainError{Code: "cancel_window_expired", Message: "item cannot be returned after 7 days from delivery"}
)

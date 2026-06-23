package exception

import (
	sharedexception "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/exception"
)

var (
	// ErrResaleOrderNotFound é retornado quando o pedido não existe, está deletado
	// (soft delete) ou não pertence ao CPF informado.
	ErrResaleOrderNotFound = &sharedexception.DomainError{
		Code:    "resale_order_not_found",
		Message: "resale order not found",
	}

	// ErrResaleOrderItemNotFound é retornado quando o item não existe dentro do pedido,
	// ou não pertence ao pedido informado.
	ErrResaleOrderItemNotFound = &sharedexception.DomainError{
		Code:    "resale_order_item_not_found",
		Message: "resale order item not found",
	}

	// ErrCancelWindowExpired é retornado quando já se passaram mais de 7 dias
	// desde a entrega (delivered_at) do item.
	ErrCancelWindowExpired = &sharedexception.DomainError{
		Code:    "cancel_window_expired",
		Message: "cancellation window of 7 days after delivery has expired",
	}
)

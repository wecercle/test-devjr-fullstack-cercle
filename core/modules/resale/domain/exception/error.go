package exception

import "errors"

var (
	ErrOrderNotFound      = errors.New("order not found or does not belong to user")
	ErrItemNotFound       = errors.New("item not found in this order")
	ErrItemNotCancellable = errors.New("item cannot be cancelled: past the 7 days limit")
)
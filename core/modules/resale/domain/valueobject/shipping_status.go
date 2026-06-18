package valueobject

import domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"

const (
	ShippingStatusLabelGenerated = "LABEL_GENERATED"
	ShippingStatusPosted         = "POSTED"
	ShippingStatusDelivered      = "DELIVERED"
	ShippingStatusReturned       = "RETURNED"
	ShippingStatusCancelled      = "CANCELLED"
)

type ShippingStatus struct {
	value string
}

func NewShippingStatus(raw string) (ShippingStatus, error) {
	switch raw {
	case ShippingStatusLabelGenerated,
		ShippingStatusPosted,
		ShippingStatusDelivered,
		ShippingStatusReturned,
		ShippingStatusCancelled:
		return ShippingStatus{value: raw}, nil
	default:
		return ShippingStatus{}, domainException.ErrItemNotEligibleForReturn
	}
}

func (s ShippingStatus) String() string { return s.value }

func (s ShippingStatus) IsReturned() bool {
	return s.value == ShippingStatusReturned
}

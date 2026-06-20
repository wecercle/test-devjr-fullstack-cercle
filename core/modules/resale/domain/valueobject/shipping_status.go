package valueobject

type ShippingStatus string

const (
	ShippingStatusLabelGenerated ShippingStatus = "LABEL_GENERATED"
	ShippingStatusPosted         ShippingStatus = "POSTED"
	ShippingStatusDelivered      ShippingStatus = "DELIVERED"
	ShippingStatusReturned       ShippingStatus = "RETURNED"
	ShippingStatusCancelled      ShippingStatus = "CANCELLED"
)

func (s ShippingStatus) String() string {
	return string(s)
}

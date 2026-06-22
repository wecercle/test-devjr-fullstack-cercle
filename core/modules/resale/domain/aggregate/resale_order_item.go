package aggregate

import (
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

const (
	ShippingStatusLabelGenerated = "LABEL_GENERATED"
	ShippingStatusPosted         = "POSTED"
	ShippingStatusDelivered      = "DELIVERED"
	ShippingStatusReturned       = "RETURNED"
	ShippingStatusCancelled      = "CANCELLED"
)

const MaxReturnWindowDays = 7

type ResaleOrderItem struct {
	id              sharedValueobject.UUID
	fkResaleOrderID sharedValueobject.UUID
	sku             string
	name            string
	quantity        int32
	amountValue     string
	shippingCode    string
	shippingStatus  string
	deliveredAt     *time.Time
}

func ReconstituteResaleOrderItem(
	id, fkResaleOrderID, sku, name string,
	quantity int32,
	amountValue, shippingCode, shippingStatus string,
	deliveredAt *time.Time,
) (*ResaleOrderItem, error) {
	itemID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, domainException.ErrInvalidItemID
	}
	orderID, err := sharedValueobject.NewUUID(fkResaleOrderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}
	return &ResaleOrderItem{
		id: itemID, fkResaleOrderID: orderID,
		sku: sku, name: name, quantity: quantity,
		amountValue: amountValue, shippingCode: shippingCode,
		shippingStatus: shippingStatus, deliveredAt: deliveredAt,
	}, nil
}

func (i *ResaleOrderItem) CanCancel() error {
	if i.shippingStatus == ShippingStatusReturned {
		return nil
	}
	if i.deliveredAt != nil {
		deadline := i.deliveredAt.AddDate(0, 0, MaxReturnWindowDays)
		if time.Now().After(deadline) {
			return domainException.ErrItemNotEligibleForReturn
		}
	}
	return nil
}

func (i *ResaleOrderItem) IsAlreadyReturned() bool {
	return i.shippingStatus == ShippingStatusReturned
}

func (i *ResaleOrderItem) ID() string              { return i.id.String() }
func (i *ResaleOrderItem) FKResaleOrderID() string { return i.fkResaleOrderID.String() }
func (i *ResaleOrderItem) SKU() string             { return i.sku }
func (i *ResaleOrderItem) Name() string            { return i.name }
func (i *ResaleOrderItem) Quantity() int32         { return i.quantity }
func (i *ResaleOrderItem) AmountValue() string     { return i.amountValue }
func (i *ResaleOrderItem) ShippingCode() string    { return i.shippingCode }
func (i *ResaleOrderItem) ShippingStatus() string  { return i.shippingStatus }
func (i *ResaleOrderItem) DeliveredAt() *time.Time { return i.deliveredAt }

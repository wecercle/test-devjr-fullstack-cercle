package aggregate

import (
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

const (
	ShippingStatusReturned = "RETURNED"
	returnPeriodDays       = 7
)

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
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

func ReconstituteResaleOrderItem(id, fkResaleOrderID, sku, name string, quantity int32, amountValue, shippingCode, shippingStatus string, deliveredAt *time.Time, createdAt, updatedAt time.Time, deletedAt *time.Time) (*ResaleOrderItem, error) {
	itemID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, domainException.ErrInvalidOrderItemID
	}
	orderID, err := sharedValueobject.NewUUID(fkResaleOrderID)
	if err != nil {
		return nil, domainException.ErrInvalidOrderID
	}

	return &ResaleOrderItem{
		id:              itemID,
		fkResaleOrderID: orderID,
		sku:             sku,
		name:            name,
		quantity:        quantity,
		amountValue:     amountValue,
		shippingCode:    shippingCode,
		shippingStatus:  shippingStatus,
		deliveredAt:     deliveredAt,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
		deletedAt:       deletedAt,
	}, nil
}

func (i *ResaleOrderItem) Return(now time.Time) error {
	if i.shippingStatus == ShippingStatusReturned {
		return nil
	}
	if i.deliveredAt == nil {
		return domainException.ErrOrderItemNotDelivered
	}
	if now.After(i.deliveredAt.AddDate(0, 0, returnPeriodDays)) {
		return domainException.ErrOrderItemReturnPeriodEnded
	}

	i.shippingStatus = ShippingStatusReturned
	i.updatedAt = now
	return nil
}

func (i *ResaleOrderItem) IsReturned() bool { return i.shippingStatus == ShippingStatusReturned }

func (i *ResaleOrderItem) ID() string              { return i.id.String() }
func (i *ResaleOrderItem) FkResaleOrderID() string { return i.fkResaleOrderID.String() }
func (i *ResaleOrderItem) SKU() string             { return i.sku }
func (i *ResaleOrderItem) Name() string            { return i.name }
func (i *ResaleOrderItem) Quantity() int32         { return i.quantity }
func (i *ResaleOrderItem) AmountValue() string     { return i.amountValue }
func (i *ResaleOrderItem) ShippingCode() string    { return i.shippingCode }
func (i *ResaleOrderItem) ShippingStatus() string  { return i.shippingStatus }
func (i *ResaleOrderItem) DeliveredAt() *time.Time { return i.deliveredAt }
func (i *ResaleOrderItem) CreatedAt() time.Time    { return i.createdAt }
func (i *ResaleOrderItem) UpdatedAt() time.Time    { return i.updatedAt }
func (i *ResaleOrderItem) DeletedAt() *time.Time   { return i.deletedAt }

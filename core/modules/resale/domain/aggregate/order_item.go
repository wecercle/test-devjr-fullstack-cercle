package aggregate

import (
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/valueobject"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

const cancelWindow = 7 * 24 * time.Hour

type OrderItem struct {
	id              sharedValueobject.UUID
	resaleOrderID   sharedValueobject.UUID
	sku             string
	name            string
	quantity        int32
	amountValue     string
	shippingCode    string
	shippingStatus  valueobject.ShippingStatus
	deliveredAt     *time.Time
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

func ReconstituteOrderItem(
	id string,
	resaleOrderID string,
	sku string,
	name string,
	quantity int32,
	amountValue string,
	shippingCode string,
	shippingStatus string,
	deliveredAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*OrderItem, error) {
	itemID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, err
	}
	orderID, err := sharedValueobject.NewUUID(resaleOrderID)
	if err != nil {
		return nil, err
	}
	status, err := valueobject.NewShippingStatus(shippingStatus)
	if err != nil {
		return nil, err
	}

	return &OrderItem{
		id:             itemID,
		resaleOrderID:  orderID,
		sku:            sku,
		name:           name,
		quantity:       quantity,
		amountValue:    amountValue,
		shippingCode:   shippingCode,
		shippingStatus: status,
		deliveredAt:    deliveredAt,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
	}, nil
}

func (i *OrderItem) RequestReturn(now time.Time) error {
	if i.shippingStatus.IsReturned() {
		return nil
	}
	if i.deliveredAt == nil {
		return domainException.ErrItemNotEligibleForReturn
	}
	if now.Sub(*i.deliveredAt) > cancelWindow {
		return domainException.ErrCancelWindowExpired
	}

	status, err := valueobject.NewShippingStatus(valueobject.ShippingStatusReturned)
	if err != nil {
		return err
	}

	i.shippingStatus = status
	i.updatedAt = now
	return nil
}

func (i *OrderItem) ID() string              { return i.id.String() }
func (i *OrderItem) ResaleOrderID() string   { return i.resaleOrderID.String() }
func (i *OrderItem) SKU() string             { return i.sku }
func (i *OrderItem) Name() string            { return i.name }
func (i *OrderItem) Quantity() int32         { return i.quantity }
func (i *OrderItem) AmountValue() string     { return i.amountValue }
func (i *OrderItem) ShippingCode() string    { return i.shippingCode }
func (i *OrderItem) ShippingStatus() string  { return i.shippingStatus.String() }
func (i *OrderItem) DeliveredAt() *time.Time { return i.deliveredAt }
func (i *OrderItem) CreatedAt() time.Time    { return i.createdAt }
func (i *OrderItem) UpdatedAt() time.Time    { return i.updatedAt }
func (i *OrderItem) DeletedAt() *time.Time   { return i.deletedAt }
func (i *OrderItem) IsReturned() bool        { return i.shippingStatus.IsReturned() }

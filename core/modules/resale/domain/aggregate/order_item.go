package aggregate

import (
	"database/sql"
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

type OrderItem struct {
	id                sharedValueobject.UUID
	fkResaleOrderID   sharedValueobject.UUID
	sku               string
	name              string
	quantity          int32
	amountValue       string
	shippingCode      sql.NullString
	shippingStatus    sql.NullString
	deliveredAt       sql.NullTime
	deletedAt         sql.NullTime
}

func ReconstituteOrderItem(
	id string,
	fkResaleOrderID string,
	sku string,
	name string,
	quantity int32,
	amountValue string,
	shippingCode sql.NullString,
	shippingStatus sql.NullString,
	deliveredAt sql.NullTime,
	deletedAt sql.NullTime,
) (*OrderItem, error) {
	itemID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, err
	}
	orderID, err := sharedValueobject.NewUUID(fkResaleOrderID)
	if err != nil {
		return nil, err
	}
	return &OrderItem{
		id:              itemID,
		fkResaleOrderID: orderID,
		sku:             sku,
		name:            name,
		quantity:        quantity,
		amountValue:     amountValue,
		shippingCode:    shippingCode,
		shippingStatus:  shippingStatus,
		deliveredAt:     deliveredAt,
		deletedAt:       deletedAt,
	}, nil
}

func (oi *OrderItem) Cancel() error {
	if oi.shippingStatus.Valid && oi.shippingStatus.String == "RETURNED" {
		return domainException.ErrItemAlreadyReturned
	}
	if oi.deliveredAt.Valid {
		daysSinceDelivery := int(time.Since(oi.deliveredAt.Time).Hours() / 24)
		if daysSinceDelivery > 7 {
			return domainException.ErrReturnWindowExpired
		}
	}
	return nil
}

func (oi *OrderItem) ID() string                     { return oi.id.String() }
func (oi *OrderItem) FkResaleOrderID() string         { return oi.fkResaleOrderID.String() }
func (oi *OrderItem) Sku() string                     { return oi.sku }
func (oi *OrderItem) Name() string                    { return oi.name }
func (oi *OrderItem) Quantity() int32                 { return oi.quantity }
func (oi *OrderItem) AmountValue() string             { return oi.amountValue }
func (oi *OrderItem) ShippingCode() sql.NullString    { return oi.shippingCode }
func (oi *OrderItem) ShippingStatus() sql.NullString  { return oi.shippingStatus }
func (oi *OrderItem) DeliveredAt() sql.NullTime       { return oi.deliveredAt }
func (oi *OrderItem) DeletedAt() sql.NullTime         { return oi.deletedAt }

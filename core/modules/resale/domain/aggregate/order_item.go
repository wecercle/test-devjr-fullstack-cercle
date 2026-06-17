package aggregate

import "time"

type OrderItem struct {
	id              string
	fkResaleOrderID string
	sku             string
	name            string
	quantity        int32
	amountValue     string
	shippingCode    string
	shippingStatus  string
	deliveredAt     *time.Time
}

func ReconstituteOrderItem(
	id string,
	fkResaleOrderID string,
	sku string,
	name string,
	quantity int32,
	amountValue string,
	shippingCode string,
	shippingStatus string,
	deliveredAt *time.Time,
) (*OrderItem, error) {

	return &OrderItem{
		id:              id,
		fkResaleOrderID: fkResaleOrderID,
		sku:             sku,
		name:            name,
		quantity:        quantity,
		amountValue:     amountValue,
		shippingCode:    shippingCode,
		shippingStatus:  shippingStatus,
		deliveredAt:     deliveredAt,
	}, nil
}

func (o *OrderItem) ID() string              { return o.id }
func (o *OrderItem) FKResaleOrderID() string { return o.fkResaleOrderID }
func (o *OrderItem) SKU() string             { return o.sku }
func (o *OrderItem) Name() string            { return o.name }
func (o *OrderItem) Quantity() int32         { return o.quantity }
func (o *OrderItem) AmountValue() string     { return o.amountValue }
func (o *OrderItem) ShippingCode() string    { return o.shippingCode }
func (o *OrderItem) ShippingStatus() string  { return o.shippingStatus }
func (o *OrderItem) DeliveredAt() *time.Time { return o.deliveredAt }

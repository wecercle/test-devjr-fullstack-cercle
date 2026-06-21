package aggregate

import (
	"time"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/valueobject"
)

type ResaleOrderItem struct {
	id              string
	fkResaleOrderID string
	sku             string
	name            string
	quantity        int32
	amountValue     string
	shippingCode    *string
	shippingStatus  valueobject.ShippingStatus
	deliveredAt     *time.Time
}

func NewResaleOrderItem(id, fkResaleOrderID, sku, name string, quantity int32, amountValue string, shippingCode *string, shippingStatus valueobject.ShippingStatus, deliveredAt *time.Time) *ResaleOrderItem {
	return &ResaleOrderItem{
		id:              id,
		fkResaleOrderID: fkResaleOrderID,
		sku:             sku,
		name:            name,
		quantity:        quantity,
		amountValue:     amountValue,
		shippingCode:    shippingCode,
		shippingStatus:  shippingStatus,
		deliveredAt:     deliveredAt,
	}
}

func (i *ResaleOrderItem) ID() string                                 { return i.id }
func (i *ResaleOrderItem) FkResaleOrderID() string                    { return i.fkResaleOrderID }
func (i *ResaleOrderItem) SKU() string                                { return i.sku }
func (i *ResaleOrderItem) Name() string                               { return i.name }
func (i *ResaleOrderItem) Quantity() int32                            { return i.quantity }
func (i *ResaleOrderItem) AmountValue() string                        { return i.amountValue }
func (i *ResaleOrderItem) ShippingCode() *string                      { return i.shippingCode }
func (i *ResaleOrderItem) ShippingStatus() valueobject.ShippingStatus { return i.shippingStatus }
func (i *ResaleOrderItem) DeliveredAt() *time.Time                    { return i.deliveredAt }

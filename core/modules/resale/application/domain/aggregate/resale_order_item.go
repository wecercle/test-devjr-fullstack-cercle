package aggregate

import (
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/exception"
)

// ResaleOrderItem representa um item de pedido de revenda.
// Os campos são privados, seguindo o mesmo padrão do agregado Retailer:
// acesso somente via getters, garantindo que o estado só mude através
// de métodos do próprio agregado (encapsulamento).
type ResaleOrderItem struct {
	id              string
	fkResaleOrderID string
	sku             string
	name            string
	quantity        int32
	amountValue     string
	shippingCode    string
	shippingStatus  string
	deliveredAt     *time.Time
	deletedAt       *time.Time
}

func ReconstituteResaleOrderItem(
	id, fkResaleOrderID, sku, name string,
	quantity int32,
	amountValue, shippingCode, shippingStatus string,
	deliveredAt *time.Time,
	deletedAt *time.Time,
) *ResaleOrderItem {
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
		deletedAt:       deletedAt,
	}
}

// CanBeCancelled valida a regra de negócio do cancelamento (devolução).
// Retorna:
//   - (true, nil)  -> o item já estava RETURNED (idempotente, handler deve responder 204)
//   - (false, nil) -> o item pode ser cancelado normalmente
//   - (false, err) -> o item NÃO pode ser cancelado (prazo de 7 dias expirado)
func (i *ResaleOrderItem) CanBeCancelled() (alreadyReturned bool, err error) {
	if i.shippingStatus == "RETURNED" {
		return true, nil
	}
	if i.deliveredAt == nil {
		// Pedido que nunca foi entregue, não tem como contar os 7 dias a partir da entrega.
		return false, domainException.ErrCancelWindowExpired
	}
	if time.Since(*i.deliveredAt) > 7*24*time.Hour {
		return false, domainException.ErrCancelWindowExpired
	}
	return false, nil
}

func (i *ResaleOrderItem) ID() string              { return i.id }
func (i *ResaleOrderItem) FkResaleOrderID() string { return i.fkResaleOrderID }
func (i *ResaleOrderItem) SKU() string             { return i.sku }
func (i *ResaleOrderItem) Name() string            { return i.name }
func (i *ResaleOrderItem) Quantity() int32         { return i.quantity }
func (i *ResaleOrderItem) AmountValue() string     { return i.amountValue }
func (i *ResaleOrderItem) ShippingCode() string    { return i.shippingCode }
func (i *ResaleOrderItem) ShippingStatus() string  { return i.shippingStatus }

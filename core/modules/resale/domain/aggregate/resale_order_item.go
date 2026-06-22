package aggregate

import "time"

// ResaleOrderItem representa o item do pedido no domínio da aplicação
type ResaleOrderItem struct {
	ID             string
	FkResaleOrderID string
	Sku            string
	Name           string
	Quantity       int
	AmountValue    float64
	ShippingCode   *string
	ShippingStatus *string
	DeliveredAt    *time.Time
}

// IsCancellable verifica a regra de negócio: se já passaram mais de 7 dias da entrega
func (r *ResaleOrderItem) IsCancellable() bool {
	if r.DeliveredAt == nil {
		return true // Se não foi entregue ainda, pode cancelar
	}
	
	// Valida se a data atual tem 7 dias ou menos de diferença da data de entrega
	deadline := r.DeliveredAt.AddDate(0, 0, 7)
	return time.Now().Before(deadline) || time.Now().Equal(deadline)
}
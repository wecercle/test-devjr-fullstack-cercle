package output

type OrderItemOutputDTO struct {
	ID              string `json:"id"`
	FKResaleOrderID string `json:"fk_resale_order_id"`
	SKU             string `json:"sku"`
	Name            string `json:"name"`
	Quantity        int32  `json:"quantity"`
	AmountValue     string `json:"amount_value"`
	ShippingCode    string `json:"shipping_code"`
	ShippingStatus  string `json:"shipping_status"`
}

type ListOrderItemOutputDTO struct {
	Items []OrderItemOutputDTO `json:"data"`
}

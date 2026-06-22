package model

import "time"

type OrderItem struct {
    ID             string    `json:"id"`
    FKResaleOrderID string    `json:"fk_resale_order_id"`
    SKU            string    `json:"sku"`
    Name           string    `json:"name"`
    Quantity       int       `json:"quantity"`
    AmountValue    string    `json:"amount_value"`
    ShippingCode   string    `json:"shipping_code"`
    ShippingStatus string    `json:"shipping_status"`
    DeliveredAt    *time.Time `json:"delivered_at,omitempty"`
    DeletedAt      *time.Time `json:"-"`
}

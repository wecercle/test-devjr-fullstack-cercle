package mapper

import (
    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/model"
)

type OrderItemMapper struct{}

func NewOrderItemMapper() *OrderItemMapper { return &OrderItemMapper{} }

func (m *OrderItemMapper) ToOutputDTO(item *model.OrderItem) output.OrderItemOutputDTO {
    return output.OrderItemOutputDTO{
        ID:             item.ID,
        FKResaleOrderID: item.FKResaleOrderID,
        SKU:            item.SKU,
        Name:           item.Name,
        Quantity:       item.Quantity,
        AmountValue:    item.AmountValue,
        ShippingCode:   item.ShippingCode,
        ShippingStatus: item.ShippingStatus,
    }
}

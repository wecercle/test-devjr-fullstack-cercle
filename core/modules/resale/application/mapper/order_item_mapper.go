package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type OrderItemMapper struct{}

func NewOrderItemMapper() *OrderItemMapper { return &OrderItemMapper{} }

func (m *OrderItemMapper) ToOutputDTO(item *aggregate.OrderItem) output.OrderItemOutputDTO {
	return output.OrderItemOutputDTO{
		ID:              item.ID(),
		FKResaleOrderID: item.FKResaleOrderID(),
		SKU:             item.SKU(),
		Name:            item.Name(),
		Quantity:        item.Quantity(),
		AmountValue:     item.AmountValue(),
		ShippingCode:    item.ShippingCode(),
		ShippingStatus:  item.ShippingStatus(),
	}
}

func (m *OrderItemMapper) ToListOutputDTO(items []*aggregate.OrderItem) output.ListOrderItemOutputDTO {
	dtoItems := make([]output.OrderItemOutputDTO, 0, len(items))
	for _, it := range items {
		dtoItems = append(dtoItems, m.ToOutputDTO(it))
	}
	return output.ListOrderItemOutputDTO{Items: dtoItems}
}

package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleOrderItemMapper struct{}

func NewResaleOrderItemMapper() *ResaleOrderItemMapper { return &ResaleOrderItemMapper{} }

func (m *ResaleOrderItemMapper) ToOutputDTO(item *aggregate.ResaleOrderItem) output.OrderItemOutputDTO {
	return output.OrderItemOutputDTO{
		ID:              item.ID(),
		FkResaleOrderID: item.FkResaleOrderID(),
		SKU:             item.SKU(),
		Name:            item.Name(),
		Quantity:        item.Quantity(),
		AmountValue:     item.AmountValue(),
		ShippingCode:    item.ShippingCode(),
		ShippingStatus:  item.ShippingStatus(),
	}
}

func (m *ResaleOrderItemMapper) ToListOutputDTO(items []*aggregate.ResaleOrderItem) output.ListOrderItemsOutputDTO {
	outputItems := make([]output.OrderItemOutputDTO, 0, len(items))
	for _, item := range items {
		outputItems = append(outputItems, m.ToOutputDTO(item))
	}
	return output.ListOrderItemsOutputDTO{Items: outputItems}
}

package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type OrderItemMapper struct{}

func NewOrderItemMapper() *OrderItemMapper { return &OrderItemMapper{} }

func (m *OrderItemMapper) ToOutputDTO(item *aggregate.OrderItem) output.OrderItemOutputDTO {
	var shippingCode *string
	if item.ShippingCode().Valid {
		sc := item.ShippingCode().String
		shippingCode = &sc
	}

	var shippingStatus *string
	if item.ShippingStatus().Valid {
		ss := item.ShippingStatus().String
		shippingStatus = &ss
	}

	return output.OrderItemOutputDTO{
		ID:              item.ID(),
		FkResaleOrderID: item.FkResaleOrderID(),
		Sku:             item.Sku(),
		Name:            item.Name(),
		Quantity:        item.Quantity(),
		AmountValue:     item.AmountValue(),
		ShippingCode:    shippingCode,
		ShippingStatus:  shippingStatus,
	}
}

func (m *OrderItemMapper) ToListOutputDTO(items []*aggregate.OrderItem) output.ListOrderItemsOutputDTO {
	data := make([]output.OrderItemOutputDTO, 0, len(items))
	for _, item := range items {
		data = append(data, m.ToOutputDTO(item))
	}
	return output.ListOrderItemsOutputDTO{Data: data}
}

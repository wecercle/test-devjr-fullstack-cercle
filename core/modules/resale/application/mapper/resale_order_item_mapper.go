package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
)

type ResaleOrderItemMapper struct{}

func NewResaleOrderItemMapper() *ResaleOrderItemMapper { return &ResaleOrderItemMapper{} }

func (m *ResaleOrderItemMapper) ToOutputDTO(item *aggregate.ResaleOrderItem) output.ResaleOrderItemOutputDTO {
	return output.ResaleOrderItemOutputDTO{
		ID: item.ID(), FKResaleOrderID: item.FKResaleOrderID(),
		SKU: item.SKU(), Name: item.Name(), Quantity: item.Quantity(),
		AmountValue: item.AmountValue(), ShippingCode: item.ShippingCode(),
		ShippingStatus: item.ShippingStatus(),
	}
}

func (m *ResaleOrderItemMapper) ToListOutputDTO(items []*aggregate.ResaleOrderItem) []output.ResaleOrderItemOutputDTO {
	result := make([]output.ResaleOrderItemOutputDTO, 0, len(items))
	for _, item := range items {
		result = append(result, m.ToOutputDTO(item))
	}
	return result
}

package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/aggregate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
)

type ResaleMapper struct{}

func NewResaleMapper() *ResaleMapper { return &ResaleMapper{} }

func (m *ResaleMapper) ToOutputDTO(item *aggregate.ResaleOrderItem) output.ResaleOrderItemOutputDTO {
	return output.ResaleOrderItemOutputDTO{
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

func (m *ResaleMapper) ToListOutputDTO(items []*aggregate.ResaleOrderItem) output.ListResaleOrderItemOutputDTO {
	dtos := make([]output.ResaleOrderItemOutputDTO, 0, len(items))
	for _, item := range items {
		dtos = append(dtos, m.ToOutputDTO(item))
	}
	return output.ListResaleOrderItemOutputDTO{Items: dtos}
}

package mapper

import (
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
)

type RetailerMapper struct{}

func NewRetailerMapper() *RetailerMapper { return &RetailerMapper{} }

// ToOutputDTO é um método do RetailerMapper que converte um objeto do tipo Retailer em um DTO de saída, mapeando os campos correspondentes.
func (m *RetailerMapper) ToOutputDTO(retailer *aggregate.Retailer) output.RetailerOutputDTO {
	return output.RetailerOutputDTO{
		ID:             retailer.ID(),
		DocumentNumber: retailer.DocumentNumber(),
		Name:           retailer.Name(),
		TradeName:      retailer.TradeName(),
		CreatedAt:      retailer.CreatedAt(),
		UpdatedAt:      retailer.UpdatedAt(),
		DeletedAt:      retailer.DeletedAt(),
	}
}

// ToListOutputDTO é um método do RetailerMapper que converte uma lista de objetos do tipo Retailer em um DTO de saída contendo uma lista de DTOs individuais, utilizando o método ToOutputDTO para cada item da lista.
func (m *RetailerMapper) ToListOutputDTO(retailers []*aggregate.Retailer) output.ListRetailerOutputDTO {
	items := make([]output.RetailerOutputDTO, 0, len(retailers))
	for _, retailer := range retailers {
		items = append(items, m.ToOutputDTO(retailer))
	}
	return output.ListRetailerOutputDTO{Items: items}
}

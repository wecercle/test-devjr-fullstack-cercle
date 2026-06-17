package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/mapper"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/query"
)

type ListRetailerUseCase struct {
	queryRepo queryRepo.RetailerQueryRepository
	mapper    *mapper.RetailerMapper
}

// NewListRetailerUseCase é um construtor para a estrutura ListRetailerUseCase, que recebe um repositório de consulta e um mapper como parâmetros e retorna uma instância da estrutura.
func NewListRetailerUseCase(queryRepo queryRepo.RetailerQueryRepository, mapper *mapper.RetailerMapper) *ListRetailerUseCase {
	return &ListRetailerUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *ListRetailerUseCase) Execute(ctx context.Context) (output.ListRetailerOutputDTO, error) {
	retailers, err := uc.queryRepo.SelectList(ctx)
	if err != nil {
		return output.ListRetailerOutputDTO{}, err
	}

	// Mapeia a lista de agregados Retailer para um DTO de saída e retorna
	return uc.mapper.ToListOutputDTO(retailers), nil
}

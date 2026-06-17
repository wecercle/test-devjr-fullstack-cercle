package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/mapper"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/query"
)

type GetRetailerByIDUseCase struct {
	queryRepo queryRepo.RetailerQueryRepository
	mapper    *mapper.RetailerMapper
}

// NewGetRetailerByIDUseCase e um construtor para a estrutura GetRetailerByIDUseCase, que recebe um repositorio de consulta e um mapper como parametros e retorna uma instancia da estrutura.
func NewGetRetailerByIDUseCase(queryRepo queryRepo.RetailerQueryRepository, mapper *mapper.RetailerMapper) *GetRetailerByIDUseCase {
	return &GetRetailerByIDUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *GetRetailerByIDUseCase) Execute(ctx context.Context, id string) (output.RetailerOutputDTO, error) {
	if err := validate.ValidateRetailerID(id); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	retailer, err := uc.queryRepo.SelectByID(ctx, id)
	if err != nil {
		return output.RetailerOutputDTO{}, err
	}

	// Mapeia o agregado Retailer para um DTO de saída e retorna
	return uc.mapper.ToOutputDTO(retailer), nil
}

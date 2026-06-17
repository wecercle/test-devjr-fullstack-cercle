package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/input"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/mapper"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/query"
)

type UpdateRetailerUseCase struct {
	queryRepo   queryRepo.RetailerQueryRepository
	commandRepo commandRepo.RetailerCommandRepository
	mapper      *mapper.RetailerMapper
}

// NewUpdateRetailerUseCase e um construtor para a estrutura UpdateRetailerUseCase, que recebe um repositorio de consulta, um repositorio de comando e um mapper como parametros e retorna uma instancia da estrutura.
func NewUpdateRetailerUseCase(queryRepo queryRepo.RetailerQueryRepository, commandRepo commandRepo.RetailerCommandRepository, mapper *mapper.RetailerMapper) *UpdateRetailerUseCase {
	return &UpdateRetailerUseCase{queryRepo: queryRepo, commandRepo: commandRepo, mapper: mapper}
}

func (uc *UpdateRetailerUseCase) Execute(ctx context.Context, id string, inputDTO input.UpdateRetailerInputDTO) (output.RetailerOutputDTO, error) {
	if err := uc.validateInput(id, inputDTO); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	retailer, err := uc.queryRepo.SelectByID(ctx, id)
	if err != nil {
		return output.RetailerOutputDTO{}, err
	}

	// Atualiza os dados do agregado Retailer com os novos valores fornecidos no DTO de entrada.
	if err := retailer.Update(inputDTO.Name, inputDTO.TradeName); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	if err := uc.commandRepo.Update(ctx, retailer); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	// Mapeia o agregado Retailer para um DTO de saída e retorna
	return uc.mapper.ToOutputDTO(retailer), nil
}

func (uc *UpdateRetailerUseCase) validateInput(id string, inputDTO input.UpdateRetailerInputDTO) error {
	return validate.ValidateRetailerRequiredFields(id, inputDTO.Name, inputDTO.TradeName)
}

package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/input"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/command"
)

type CreateRetailerUseCase struct {
	commandRepo commandRepo.RetailerCommandRepository
	mapper      *mapper.RetailerMapper
}

// NewCreateRetailerUseCase e um construtor para a estrutura CreateRetailerUseCase, que recebe um repositorio de comando e um mapper como parametros e retorna uma instancia da estrutura.
func NewCreateRetailerUseCase(commandRepo commandRepo.RetailerCommandRepository, mapper *mapper.RetailerMapper) *CreateRetailerUseCase {
	return &CreateRetailerUseCase{commandRepo: commandRepo, mapper: mapper}
}

func (uc *CreateRetailerUseCase) Execute(ctx context.Context, inputDTO input.CreateRetailerInputDTO) (output.RetailerOutputDTO, error) {
	// Valida os dados de entrada
	if err := uc.validateInput(inputDTO); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	// gera um novo ID para o varejista e cria uma nova instancia do agregado Retailer
	id := uuid.New().String()
	retailer, err := aggregate.NewRetailer(id, inputDTO.DocumentNumber, inputDTO.Name, inputDTO.TradeName)
	if err != nil {
		return output.RetailerOutputDTO{}, err
	}

	if err := uc.commandRepo.Insert(ctx, retailer); err != nil {
		return output.RetailerOutputDTO{}, err
	}

	// Mapeia o agregado Retailer para um DTO de saída e retorna
	return uc.mapper.ToOutputDTO(retailer), nil
}

func (uc *CreateRetailerUseCase) validateInput(inputDTO input.CreateRetailerInputDTO) error {
	return validate.ValidateRetailerRequiredFieldsForCreate(inputDTO.DocumentNumber, inputDTO.Name, inputDTO.TradeName)
}

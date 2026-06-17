package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/validate"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/repository/query"
)

type DeleteRetailerUseCase struct {
	queryRepo   queryRepo.RetailerQueryRepository
	commandRepo commandRepo.RetailerCommandRepository
}

// NewDeleteRetailerUseCase e um construtor para a estrutura DeleteRetailerUseCase, que recebe um repositorio de consulta e um repositorio de comando como parametros e retorna uma instancia da estrutura.
func NewDeleteRetailerUseCase(queryRepo queryRepo.RetailerQueryRepository, commandRepo commandRepo.RetailerCommandRepository) *DeleteRetailerUseCase {
	return &DeleteRetailerUseCase{queryRepo: queryRepo, commandRepo: commandRepo}
}

func (uc *DeleteRetailerUseCase) Execute(ctx context.Context, id string) error {
	if err := validate.ValidateRetailerID(id); err != nil {
		return err
	}

	_, err := uc.queryRepo.SelectByID(ctx, id)
	if err != nil {
		// Já foi deletado ou não existe - retorna sucesso (idempotente)
		return nil
	}

	if err := uc.commandRepo.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}

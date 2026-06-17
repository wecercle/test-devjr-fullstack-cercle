package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type CancelOrderItemUseCase struct {
	queryRepo   queryRepo.OrderItemQueryRepository
	commandRepo commandRepo.OrderItemCommandRepository
}

func NewCancelOrderItemUseCase(queryRepo queryRepo.OrderItemQueryRepository, commandRepo commandRepo.OrderItemCommandRepository) *CancelOrderItemUseCase {
	return &CancelOrderItemUseCase{queryRepo: queryRepo, commandRepo: commandRepo}
}

func (uc *CancelOrderItemUseCase) Execute(ctx context.Context, cpf string, orderID string, itemID string) error {
	if err := validate.ValidateCPF(cpf); err != nil {
		return err
	}
	if err := validate.ValidateOrderID(orderID); err != nil {
		return err
	}
	if err := validate.ValidateItemID(itemID); err != nil {
		return err
	}

	item, err := uc.queryRepo.SelectItemByID(ctx, cpf, orderID, itemID)
	if err != nil {
		return err
	}

	if err := item.Cancel(); err != nil {
		if err == domainException.ErrItemAlreadyReturned {
			return nil
		}
		return err
	}

	return uc.commandRepo.CancelItem(ctx, orderID, itemID)
}

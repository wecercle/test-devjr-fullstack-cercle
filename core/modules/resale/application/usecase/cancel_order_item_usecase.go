package usecase

import (
	"context"
	"time"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type CancelOrderItemUseCase struct {
	queryRepo   queryRepo.ResaleOrderQueryRepository
	commandRepo commandRepo.ResaleOrderCommandRepository
}

func NewCancelOrderItemUseCase(queryRepo queryRepo.ResaleOrderQueryRepository, commandRepo commandRepo.ResaleOrderCommandRepository) *CancelOrderItemUseCase {
	return &CancelOrderItemUseCase{queryRepo: queryRepo, commandRepo: commandRepo}
}

func (uc *CancelOrderItemUseCase) Execute(ctx context.Context, cpf, orderID, itemID string) error {
	if err := validate.ValidateOrderID(orderID); err != nil {
		return err
	}
	if err := validate.ValidateOrderItemID(itemID); err != nil {
		return err
	}

	item, err := uc.queryRepo.SelectItemByCPFOrderIDAndItemID(ctx, cpf, orderID, itemID)
	if err != nil {
		return err
	}

	if item.IsReturned() {
		return nil
	}
	if err := item.Return(time.Now()); err != nil {
		return err
	}

	return uc.commandRepo.UpdateItemShippingStatus(ctx, item)
}

package usecase

import (
	"context"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

type CancelOrderItemUseCase struct {
	queryRepo   queryRepo.ResaleOrderItemQueryRepository
	commandRepo commandRepo.ResaleOrderItemCommandRepository
}

func NewCancelOrderItemUseCase(q queryRepo.ResaleOrderItemQueryRepository, c commandRepo.ResaleOrderItemCommandRepository) *CancelOrderItemUseCase {
	return &CancelOrderItemUseCase{queryRepo: q, commandRepo: c}
}

// Retorna (true, nil) se idempotente, (false, nil) se cancelado, (false, error) se erro.
func (uc *CancelOrderItemUseCase) Execute(ctx context.Context, cpf, orderID, itemID string) (bool, error) {
	if _, err := sharedValueobject.NewUUID(orderID); err != nil {
		return false, domainException.ErrInvalidOrderID
	}
	if _, err := sharedValueobject.NewUUID(itemID); err != nil {
		return false, domainException.ErrInvalidItemID
	}
	item, err := uc.queryRepo.SelectByIDAndOrderIDAndCPF(ctx, itemID, orderID, cpf)
	if err != nil {
		return false, err
	}
	if item.IsAlreadyReturned() {
		return true, nil
	}
	if err = item.CanCancel(); err != nil {
		return false, err
	}
	if err = uc.commandRepo.UpdateShippingStatus(ctx, itemID, orderID, "RETURNED"); err != nil {
		return false, err
	}
	return false, nil
}

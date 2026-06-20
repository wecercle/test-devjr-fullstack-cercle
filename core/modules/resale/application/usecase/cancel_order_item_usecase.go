package usecase

import (
	"context"
	"time"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/valueobject"
)

const returnPeriodDays = 7

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

	exists, err := uc.queryRepo.ExistsOrderByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return err
	}
	if !exists {
		return domainException.ErrOrderNotFound
	}

	item, err := uc.queryRepo.SelectOrderItemByOrderIDAndItemID(ctx, orderID, itemID)
	if err != nil {
		return err
	}

	if item.ShippingStatus() == valueobject.ShippingStatusReturned {
		return nil
	}
	if item.ShippingStatus() != valueobject.ShippingStatusDelivered {
		return domainException.ErrOrderItemNotReturnable
	}
	if item.DeliveredAt() == nil {
		return domainException.ErrOrderItemNotReturnable
	}
	if time.Now().UTC().After(item.DeliveredAt().UTC().AddDate(0, 0, returnPeriodDays)) {
		return domainException.ErrOrderItemReturnExpired
	}

	updated, err := uc.commandRepo.MarkOrderItemAsReturned(ctx, orderID, itemID)
	if err != nil {
		return err
	}
	if !updated {
		return domainException.ErrOrderItemNotReturnable
	}

	return nil
}

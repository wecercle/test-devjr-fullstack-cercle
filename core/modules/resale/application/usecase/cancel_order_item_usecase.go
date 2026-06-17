package usecase

import (
	"context"
	"time"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

const ReturnedStatus = "RETURNED"

type CancelOrderItemUseCase struct {
	queryRepo   queryRepo.OrderQueryRepository
	commandRepo commandRepo.OrderCommandRepository
}

func NewCancelOrderItemUseCase(queryRepo queryRepo.OrderQueryRepository, commandRepo commandRepo.OrderCommandRepository) *CancelOrderItemUseCase {
	return &CancelOrderItemUseCase{queryRepo: queryRepo, commandRepo: commandRepo}
}

func (uc *CancelOrderItemUseCase) Execute(ctx context.Context, cpf string, orderID string, itemID string) error {
	items, err := uc.queryRepo.SelectOrderItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return domainException.ErrOrderNotFound
	}

	var targetItemFound bool
	var targetItem *aggregate.OrderItem
	for _, it := range items {
		if it.ID() == itemID {
			targetItemFound = true
			targetItem = it
			break
		}
	}
	if !targetItemFound {
		return domainException.ErrOrderItemNotFound
	}

	// If already returned, idempotent
	if targetItem.ShippingStatus() == ReturnedStatus {
		return domainException.ErrOrderItemAlreadyReturned
	}

	// Must have delivered_at and be within 7 days
	deliveredAt := targetItem.DeliveredAt()
	if deliveredAt == nil {
		return domainException.ErrOrderItemCancelWindowExpired
	}
	if time.Since(*deliveredAt) > 7*24*time.Hour {
		return domainException.ErrOrderItemCancelWindowExpired
	}

	if err := uc.commandRepo.UpdateOrderItemShippingStatus(ctx, orderID, itemID, ReturnedStatus); err != nil {
		return err
	}
	return nil
}

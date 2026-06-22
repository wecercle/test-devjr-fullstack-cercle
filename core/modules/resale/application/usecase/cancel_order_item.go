package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository"
)

type CancelOrderItemUseCase struct {
	queryRepo   repository.ResaleQueryRepository
	commandRepo repository.ResaleCommandRepository
}

func NewCancelOrderItemUseCase(queryRepo repository.ResaleQueryRepository, commandRepo repository.ResaleCommandRepository) *CancelOrderItemUseCase {
	return &CancelOrderItemUseCase{
		queryRepo:   queryRepo,
		commandRepo: commandRepo,
	}
}

func (uc *CancelOrderItemUseCase) Execute(ctx context.Context, cpf, orderID, itemID string) error {
	if err := uuid.Validate(orderID); err != nil {
		return fmt.Errorf("invalid order_id: %w", err)
	}
	if err := uuid.Validate(itemID); err != nil {
		return fmt.Errorf("invalid item_id: %w", err)
	}

	item, err := uc.queryRepo.GetOrderItemForCancellation(ctx, cpf, orderID, itemID)
	if err != nil || item == nil {
		return exception.ErrOrderNotFound
	}

	if item.ShippingStatus != nil && *item.ShippingStatus == "RETURNED" {
		return nil
	}

	if !item.IsCancellable() {
		return exception.ErrItemNotCancellable
	}

	return uc.commandRepo.CancelOrderItem(ctx, itemID)
}
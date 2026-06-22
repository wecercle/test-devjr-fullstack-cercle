package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository"
)

type GetOrderItemsUseCase struct {
	queryRepo repository.ResaleQueryRepository
}

func NewGetOrderItemsUseCase(queryRepo repository.ResaleQueryRepository) *GetOrderItemsUseCase {
	return &GetOrderItemsUseCase{queryRepo: queryRepo}
}

func (uc *GetOrderItemsUseCase) Execute(ctx context.Context, cpf, orderID string) ([]dto.OrderItemResponse, error) {
	if err := uuid.Validate(orderID); err != nil {
		return nil, fmt.Errorf("invalid order_id: %w", err)
	}

	items, err := uc.queryRepo.GetOrderItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return nil, exception.ErrOrderNotFound
	}

	if len(items) == 0 {
		return nil, exception.ErrOrderNotFound
	}

	var response []dto.OrderItemResponse
	for _, item := range items {
		shippingCode := ""
		if item.ShippingCode != nil {
			shippingCode = *item.ShippingCode
		}

		shippingStatus := ""
		if item.ShippingStatus != nil {
			shippingStatus = *item.ShippingStatus
		}

		response = append(response, dto.OrderItemResponse{
			ID:             item.ID,
			FkResaleOrderID: item.FkResaleOrderID,
			Sku:            item.Sku,
			Name:           item.Name,
			Quantity:       item.Quantity,
			AmountValue:    fmt.Sprintf("%.2f", item.AmountValue),
			ShippingCode:   shippingCode,
			ShippingStatus: shippingStatus,
		})
	}

	return response, nil
}
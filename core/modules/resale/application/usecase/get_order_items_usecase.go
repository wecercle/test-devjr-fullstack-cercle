package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type GetOrderItemsUseCase struct {
	queryRepo queryRepo.OrderItemQueryRepository
	mapper    *mapper.OrderItemMapper
}

func NewGetOrderItemsUseCase(queryRepo queryRepo.OrderItemQueryRepository, mapper *mapper.OrderItemMapper) *GetOrderItemsUseCase {
	return &GetOrderItemsUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *GetOrderItemsUseCase) Execute(ctx context.Context, cpf string, orderID string) (output.ListOrderItemsOutputDTO, error) {
	if err := validate.ValidateCPF(cpf); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}
	if err := validate.ValidateOrderID(orderID); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	items, err := uc.queryRepo.SelectItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	if len(items) == 0 {
		return output.ListOrderItemsOutputDTO{Data: []output.OrderItemOutputDTO{}}, nil
	}

	return uc.mapper.ToListOutputDTO(items), nil
}

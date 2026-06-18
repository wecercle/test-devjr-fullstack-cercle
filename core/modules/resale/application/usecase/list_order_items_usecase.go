package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ListOrderItemsUseCase struct {
	queryRepo queryRepo.ResaleOrderQueryRepository
	mapper    *mapper.ResaleOrderItemMapper
}

func NewListOrderItemsUseCase(queryRepo queryRepo.ResaleOrderQueryRepository, mapper *mapper.ResaleOrderItemMapper) *ListOrderItemsUseCase {
	return &ListOrderItemsUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *ListOrderItemsUseCase) Execute(ctx context.Context, cpf, orderID string) (output.ListOrderItemsOutputDTO, error) {
	if err := validate.ValidateOrderID(orderID); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}
	if err := uc.queryRepo.SelectOrderByCPFAndOrderID(ctx, cpf, orderID); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	items, err := uc.queryRepo.SelectItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	return uc.mapper.ToListOutputDTO(items), nil
}

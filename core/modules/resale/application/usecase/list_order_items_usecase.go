package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ListOrderItemsUseCase struct {
	queryRepo queryRepo.OrderQueryRepository
	mapper    *mapper.OrderItemMapper
}

func NewListOrderItemsUseCase(queryRepo queryRepo.OrderQueryRepository, mapper *mapper.OrderItemMapper) *ListOrderItemsUseCase {
	return &ListOrderItemsUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *ListOrderItemsUseCase) Execute(ctx context.Context, cpf string, orderID string) (output.ListOrderItemOutputDTO, error) {
	items, err := uc.queryRepo.SelectOrderItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return output.ListOrderItemOutputDTO{}, err
	}
	if len(items) == 0 {
		return output.ListOrderItemOutputDTO{}, domainException.ErrOrderNotFound
	}
	return uc.mapper.ToListOutputDTO(items), nil
}

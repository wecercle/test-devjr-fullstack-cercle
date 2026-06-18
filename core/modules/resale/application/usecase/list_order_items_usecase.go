package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/validate"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ListOrderItemsUseCase struct {
	queryRepo queryRepo.ResaleQueryRepository
	mapper    *mapper.OrderItemMapper
}

func NewListOrderItemsUseCase(queryRepo queryRepo.ResaleQueryRepository, mapper *mapper.OrderItemMapper) *ListOrderItemsUseCase {
	return &ListOrderItemsUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *ListOrderItemsUseCase) Execute(ctx context.Context, cpf string, orderID string) (output.ListOrderItemsOutputDTO, error) {
	if err := validate.ValidateUserDocumentNumber(cpf); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}
	if err := validate.ValidateOrderID(orderID); err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	exists, err := uc.queryRepo.OrderExistsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}
	if !exists {
		return output.ListOrderItemsOutputDTO{}, domainException.ErrOrderNotFound
	}

	items, err := uc.queryRepo.SelectOrderItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return output.ListOrderItemsOutputDTO{}, err
	}

	return uc.mapper.ToListOutputDTO(items), nil
}

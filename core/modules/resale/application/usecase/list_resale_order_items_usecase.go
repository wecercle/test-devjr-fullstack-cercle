package usecase

import (
	"context"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/exception"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/repository/query"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
)

// ListResaleOrderItemsUseCase implementa o GET /v1/app/users/:cpf/orders/:order_id/items
type ListResaleOrderItemsUseCase struct {
	queryRepo queryRepo.ResaleQueryRepository
	mapper    *mapper.ResaleMapper
}

func NewListResaleOrderItemsUseCase(queryRepo queryRepo.ResaleQueryRepository, mapper *mapper.ResaleMapper) *ListResaleOrderItemsUseCase {
	return &ListResaleOrderItemsUseCase{queryRepo: queryRepo, mapper: mapper}
}

func (uc *ListResaleOrderItemsUseCase) Execute(ctx context.Context, orderID, cpf string) (output.ListResaleOrderItemOutputDTO, error) {
	items, err := uc.queryRepo.SelectItemsByOrderAndCPF(ctx, orderID, cpf)
	if err != nil {
		return output.ListResaleOrderItemOutputDTO{}, err
	}

	if len(items) == 0 {
		return output.ListResaleOrderItemOutputDTO{}, domainException.ErrResaleOrderNotFound
	}

	return uc.mapper.ToListOutputDTO(items), nil
}

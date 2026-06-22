package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

type GetOrderItemsUseCase struct {
	queryRepo queryRepo.ResaleOrderItemQueryRepository
	mapper    *mapper.ResaleOrderItemMapper
}

func NewGetOrderItemsUseCase(q queryRepo.ResaleOrderItemQueryRepository, m *mapper.ResaleOrderItemMapper) *GetOrderItemsUseCase {
	return &GetOrderItemsUseCase{queryRepo: q, mapper: m}
}

func (uc *GetOrderItemsUseCase) Execute(ctx context.Context, cpf, orderID string) ([]output.ResaleOrderItemOutputDTO, error) {
	if _, err := sharedValueobject.NewUUID(orderID); err != nil {
		return nil, domainException.ErrInvalidOrderID
	}
	items, err := uc.queryRepo.SelectByOrderIDAndCPF(ctx, orderID, cpf)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, domainException.ErrOrderNotFound
	}
	return uc.mapper.ToListOutputDTO(items), nil
}

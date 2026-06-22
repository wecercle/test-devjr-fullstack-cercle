package usecase

import (
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
    "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/model"
    queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ListOrderItemsUseCase struct {
    queryRepo queryRepo.OrderItemQueryRepository
    mapper    *mapper.OrderItemMapper
}

func NewListOrderItemsUseCase(q queryRepo.OrderItemQueryRepository, m *mapper.OrderItemMapper) *ListOrderItemsUseCase {
    return &ListOrderItemsUseCase{queryRepo: q, mapper: m}
}

// validates CPF (only digits, length 11) and UUIDs
func validateCPF(cpf string) error {
    if len(cpf) != 11 {
        return errors.New("invalid CPF length")
    }
    for _, r := range cpf {
        if r < '0' || r > '9' {
            return errors.New("CPF must contain only digits")
        }
    }
    return nil
}

func (uc *ListOrderItemsUseCase) Execute(ctx context.Context, cpf string, orderID string) ([]output.OrderItemOutputDTO, error) {
    if err := validateCPF(cpf); err != nil {
        return nil, err
    }
    if _, err := uuid.Parse(orderID); err != nil {
        return nil, errors.New("invalid order_id UUID")
    }

    items, err := uc.queryRepo.ListByCPFAndOrderID(ctx, cpf, orderID)
    if err != nil {
        return nil, err
    }
    var result []output.OrderItemOutputDTO
    for _, it := range items {
        if it.DeletedAt != nil {
            continue // ignore soft‑deleted items
        }
        result = append(result, uc.mapper.ToOutputDTO(it))
    }
    return result, nil
}

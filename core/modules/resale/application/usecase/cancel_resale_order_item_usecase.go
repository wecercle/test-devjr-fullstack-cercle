package usecase

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/exception"
	commandRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/repository/command"
	queryRepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/repository/query"
)

// CancelResaleOrderItemUseCase
// PUT /v1/app/users/:cpf/orders/:order_id/items/:item_id/cancel
type CancelResaleOrderItemUseCase struct {
	queryRepo   queryRepo.ResaleQueryRepository
	commandRepo commandRepo.ResaleCommandRepository
}

func NewCancelResaleOrderItemUseCase(queryRepo queryRepo.ResaleQueryRepository, commandRepo commandRepo.ResaleCommandRepository) *CancelResaleOrderItemUseCase {
	return &CancelResaleOrderItemUseCase{queryRepo: queryRepo, commandRepo: commandRepo}
}

func (uc *CancelResaleOrderItemUseCase) Execute(ctx context.Context, orderID, itemID, cpf string) error {
	// 1. Busca todos os itens do pedido. A query já garante que o pedido
	//    pertence ao CPF informado e que nem o pedido nem os itens estão
	//    soft-deletados.
	items, err := uc.queryRepo.SelectItemsByOrderAndCPF(ctx, orderID, cpf)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return domainException.ErrResaleOrderNotFound
	}

	// 2. Localiza o item específico dentro da lista (filtragem em memória,
	//    decisão tomada para não precisar de uma query SQL adicional).
	var target *aggregate.ResaleOrderItem
	for _, item := range items {
		if item.ID() == itemID {
			target = item
			break
		}
	}
	if target == nil {
		return domainException.ErrResaleOrderItemNotFound
	}

	// 3. Valida a regra de negócio: já está RETURNED? passou de 7 dias?
	alreadyReturned, err := target.CanBeCancelled()
	if err != nil {
		return err // ErrCancelWindowExpired -> handler responde 400
	}
	if alreadyReturned {
		return nil // idempotente -> handler responde 204 do mesmo jeito
	}

	// 4. Persiste a mudança de status. :execrows retorna quantas linhas
	//    foram afetadas — usamos isso como segunda camada de segurança,
	//    caso o item tenha sido alterado entre o passo 1 e este passo 4.
	rows, err := uc.commandRepo.UpdateShippingStatus(ctx, orderID, itemID, "RETURNED")
	if err != nil {
		return err
	}
	if rows == 0 {
		return domainException.ErrResaleOrderItemNotFound
	}

	return nil
}

package command

import (
	"context"
)

type ResaleCommandRepository interface {
	// UpdateShippingStatus atualiza o shipping_status de um item específico,
	// filtrando também pelo pedido (orderID) para garantir que o item realmente
	// pertence àquele pedido.
	//
	// Retorna o número de linhas afetadas (rowsAffected). Se for 0, significa que
	// o item não existe ou não pertence ao pedido informado — usado pelo use case
	// para decidir se deve retornar ErrResaleOrderItemNotFound.
	UpdateShippingStatus(ctx context.Context, orderID, itemID, status string) (int64, error)
}

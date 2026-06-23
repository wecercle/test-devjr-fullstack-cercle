package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/aggregate"
)

// Existe apenas um método porque a estratégia adotada foi reaproveitar a
// listagem (SelectItemsByOrderAndCPF) tanto para o GET quanto para localizar
// um item específico antes de cancelá-lo (filtragem feita em memória no use case).
type ResaleQueryRepository interface {
	SelectItemsByOrderAndCPF(ctx context.Context, orderID, cpf string) ([]*aggregate.ResaleOrderItem, error)
}

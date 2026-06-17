package query

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
)

// RetailerQueryRepository define a interface de repositório de consulta para o agregado Retailer.
// Light CQRS: este repositório é responsável apenas por operações de leitura (select) no banco de dados.
type RetailerQueryRepository interface {
	SelectList(ctx context.Context) ([]*aggregate.Retailer, error)
	SelectByID(ctx context.Context, id string) (*aggregate.Retailer, error)
}

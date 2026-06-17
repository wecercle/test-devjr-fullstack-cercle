package command

import (
	"context"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
)

// RetailerCommandRepository definee a interface de repositório de comando para o agregado Retailer.
// Light CQRS: este repositório é responsável apenas por operações de escrita (insert, update, delete) no banco de dados.
type RetailerCommandRepository interface {
	Insert(ctx context.Context, retailer *aggregate.Retailer) error
	Update(ctx context.Context, retailer *aggregate.Retailer) error
	SoftDelete(ctx context.Context, id string) error
}

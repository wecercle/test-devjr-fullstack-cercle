package pgcommand

import (
	"context"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository"
)

type ResaleCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleCommandRepository(querier *databaseQuery.Queries) repository.ResaleCommandRepository {
	return &ResaleCommandRepository{querier: querier}
}

func (r *ResaleCommandRepository) CancelOrderItem(ctx context.Context, itemID string) error {
	
	uid, err := uuid.Parse(itemID)
	if err != nil {
		return err
	}
	return r.querier.CancelOrderItem(ctx, uid)
}
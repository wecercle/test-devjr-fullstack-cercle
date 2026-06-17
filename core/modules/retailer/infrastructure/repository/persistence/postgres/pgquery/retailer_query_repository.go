package pgquery

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
	sharedInfrastructure "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/infrastructure"
)

type RetailerQueryRepository struct {
	querier *databaseQuery.Queries
}

// NewRetailerQueryRepository é um construtor para a estrutura RetailerQueryRepository, que recebe um querier como parâmetro e retorna uma instância da estrutura.
func NewRetailerQueryRepository(querier *databaseQuery.Queries) *RetailerQueryRepository {
	return &RetailerQueryRepository{querier: querier}
}

// SelectList recupera uma lista de varejistas do banco de dados, reconstituindo cada um como um agregado Retailer. Retorna um slice de ponteiros para Retailer e um erro, se houver.
func (r *RetailerQueryRepository) SelectList(ctx context.Context) ([]*aggregate.Retailer, error) {
	rows, err := r.querier.SelectRetailerList(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*aggregate.Retailer, 0, len(rows))
	for _, row := range rows {
		retailer, err := aggregate.ReconstituteRetailer(
			row.ID.String(),
			row.DocumentNumber,
			row.Name,
			row.TradeName,
			row.CreatedAt,
			row.UpdatedAt,
			sharedInfrastructure.NullTimeToPointer(row.DeletedAt),
		)
		if err != nil {
			return nil, err
		}
		items = append(items, retailer)
	}

	return items, nil
}

// SelectByID recupera um varejista do banco de dados pelo ID, reconstituindo-o como um agregado Retailer. Retorna um ponteiro para Retailer e um erro, se houver.
func (r *RetailerQueryRepository) SelectByID(ctx context.Context, id string) (*aggregate.Retailer, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	row, err := r.querier.SelectRetailerByID(ctx, parsedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainException.ErrRetailerNotFound
		}
		return nil, err
	}

	retailer, err := aggregate.ReconstituteRetailer(
		row.ID.String(),
		row.DocumentNumber,
		row.Name,
		row.TradeName,
		row.CreatedAt,
		row.UpdatedAt,
		sharedInfrastructure.NullTimeToPointer(row.DeletedAt),
	)
	if err != nil {
		return nil, err
	}

	return retailer, nil
}

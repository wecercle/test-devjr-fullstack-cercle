package pgcommand

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
)

type RetailerCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewRetailerCommandRepository(querier *databaseQuery.Queries) *RetailerCommandRepository {
	return &RetailerCommandRepository{querier: querier}
}

// Insert insere um novo varejista no banco de dados, utilizando os dados do agregado Retailer. Retorna um erro, se houver.
func (r *RetailerCommandRepository) Insert(ctx context.Context, retailer *aggregate.Retailer) error {
	var deletedAt sql.NullTime
	if retailer.DeletedAt() != nil {
		deletedAt = sql.NullTime{Time: *retailer.DeletedAt(), Valid: true}
	}

	parsedID, err := uuid.Parse(retailer.ID())
	if err != nil {
		return err
	}

	err = r.querier.InsertRetailer(ctx, databaseQuery.InsertRetailerParams{
		ID:             parsedID,
		DocumentNumber: retailer.DocumentNumber(),
		Name:           retailer.Name(),
		TradeName:      retailer.TradeName(),
		CreatedAt:      retailer.CreatedAt(),
		UpdatedAt:      retailer.UpdatedAt(),
		DeletedAt:      deletedAt,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// unique_violation
			if pgErr.Code == "23505" {
				return domainException.ErrRetailerDocumentNumberExists
			}
		}
		return err
	}
	return nil
}

// Update atualiza os dados de um varejista existente no banco de dados, utilizando os dados do agregado Retailer. Retorna um erro, se houver.
func (r *RetailerCommandRepository) Update(ctx context.Context, retailer *aggregate.Retailer) error {
	parsedID, err := uuid.Parse(retailer.ID())
	if err != nil {
		return err
	}

	rowsAffected, err := r.querier.UpdateRetailer(ctx, databaseQuery.UpdateRetailerParams{
		Name:      retailer.Name(),
		TradeName: retailer.TradeName(),
		UpdatedAt: retailer.UpdatedAt(),
		ID:        parsedID,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domainException.ErrRetailerNotFound
	}
	return nil
}

// SoftDelete realiza uma exclusão lógica de um varejista no banco de dados, utilizando o ID do varejista. Retorna um erro, se houver.
func (r *RetailerCommandRepository) SoftDelete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	rowsAffected, err := r.querier.SoftDeleteRetailer(ctx, parsedID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domainException.ErrRetailerNotFound
	}
	return nil
}

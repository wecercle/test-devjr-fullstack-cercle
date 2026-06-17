package retailer

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/infrastructure/repository/persistence/postgres/pgcommand"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/infrastructure/repository/persistence/postgres/pgquery"
	retailerhttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/presentation/http"
)

// Container agrupa as dependências do módulo Retailer
type Container struct {
	Handler *retailerhttp.Handler
}

// Setup inicializa o módulo Retailer com dependency injection manual
func Setup(querier *databaseQuery.Queries) *Container {
	// Infrastructure Layer
	queryRepo := pgquery.NewRetailerQueryRepository(querier)
	commandRepo := pgcommand.NewRetailerCommandRepository(querier)

	// Application Layer - Use Cases
	retailerMapper := mapper.NewRetailerMapper()
	createUseCase := usecase.NewCreateRetailerUseCase(commandRepo, retailerMapper)
	updateUseCase := usecase.NewUpdateRetailerUseCase(queryRepo, commandRepo, retailerMapper)
	listUseCase := usecase.NewListRetailerUseCase(queryRepo, retailerMapper)
	getByIDUseCase := usecase.NewGetRetailerByIDUseCase(queryRepo, retailerMapper)
	deleteUseCase := usecase.NewDeleteRetailerUseCase(queryRepo, commandRepo)

	// Presentation Layer - HTTP Handler
	handler := retailerhttp.NewHandler(createUseCase, updateUseCase, listUseCase, getByIDUseCase, deleteUseCase)

	return &Container{Handler: handler}
}

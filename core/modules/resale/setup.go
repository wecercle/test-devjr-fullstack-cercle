package resale

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgcommand"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgquery"
	resalehttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/presentation/http"
)

// Container agrupa as dependências do módulo Resale
type Container struct {
	Handler *resalehttp.Handler
}

// Setup inicializa o módulo Resale com dependency injection manual
func Setup(querier *databaseQuery.Queries) *Container {
	// Infrastructure Layer
	queryRepo := pgquery.NewResaleOrderQueryRepository(querier)
	commandRepo := pgcommand.NewResaleOrderCommandRepository(querier)

	// Application Layer - Use Cases
	resaleOrderItemMapper := mapper.NewResaleOrderItemMapper()
	listUseCase := usecase.NewListOrderItemsByCPFAndOrderIDUseCase(queryRepo, resaleOrderItemMapper)
	cancelUseCase := usecase.NewCancelOrderItemUseCase(queryRepo, commandRepo)

	// Presentation Layer - HTTP Handler
	handler := resalehttp.NewHandler(listUseCase, cancelUseCase)

	return &Container{Handler: handler}
}

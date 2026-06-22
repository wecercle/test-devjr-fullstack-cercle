package resale

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgcommand"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgquery"
	resalehttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/presentation/http"
)

type Container struct {
	Handler *resalehttp.Handler
}

func Setup(querier *databaseQuery.Queries) *Container {
	// Instancia da Infraestrutura
	queryRepo := pgquery.NewResaleQueryRepository(querier)
	commandRepo := pgcommand.NewResaleCommandRepository(querier)

	//  Instancia dos Casos de Uso
	getOrderItemsUC := usecase.NewGetOrderItemsUseCase(queryRepo)
	cancelOrderItemUC := usecase.NewCancelOrderItemUseCase(queryRepo, commandRepo)

	// Instancia o Handler
	handler := resalehttp.NewHandler(getOrderItemsUC, cancelOrderItemUC)

	return &Container{Handler: handler}
}
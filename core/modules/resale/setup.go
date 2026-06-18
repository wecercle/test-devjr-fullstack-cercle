package resale

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgcommand"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgquery"
	resalehttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/presentation/http"
)

type Container struct {
	Handler *resalehttp.Handler
}

func Setup(querier *databaseQuery.Queries) *Container {
	queryRepo := pgquery.NewResaleQueryRepository(querier)
	commandRepo := pgcommand.NewResaleCommandRepository(querier)

	orderItemMapper := mapper.NewOrderItemMapper()
	listOrderItemsUseCase := usecase.NewListOrderItemsUseCase(queryRepo, orderItemMapper)
	cancelOrderItemUseCase := usecase.NewCancelOrderItemUseCase(queryRepo, commandRepo)

	handler := resalehttp.NewHandler(listOrderItemsUseCase, cancelOrderItemUseCase)

	return &Container{Handler: handler}
}

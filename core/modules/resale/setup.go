package resale

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgcommand"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/infrastructure/repository/persistence/postgres/pgquery"
	resalehttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/presentation/http"
)

type Container struct{ Handler *resalehttp.Handler }

func Setup(querier *databaseQuery.Queries) *Container {
	queryRepo := pgquery.NewResaleOrderItemQueryRepository(querier)
	commandRepo := pgcommand.NewResaleOrderItemCommandRepository(querier)
	itemMapper := mapper.NewResaleOrderItemMapper()
	getUC := usecase.NewGetOrderItemsUseCase(queryRepo, itemMapper)
	cancelUC := usecase.NewCancelOrderItemUseCase(queryRepo, commandRepo)
	handler := resalehttp.NewHandler(getUC, cancelUC)
	return &Container{Handler: handler}
}

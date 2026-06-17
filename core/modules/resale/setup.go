package resale

import (
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	resalehttp "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/presentation/http"
)

// Container agrupa as dependências do módulo Resale
type Container struct {
	Handler *resalehttp.Handler
}

// Setup inicializa o módulo Resale com dependency injection manual
func Setup(querier *databaseQuery.Queries) *Container {
	_ = querier

	// Dummy handler while module implementation is pending.
	handler := resalehttp.NewHandler()

	return &Container{Handler: handler}
}

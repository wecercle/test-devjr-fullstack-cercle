package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/database/connection"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer"
)

type Routes struct{}

func (Routes) SetupRouterV1(router *gin.Engine) {
	q := connection.Querier()

	resaleContainer := resale.Setup(q)
	retailerContainer := retailer.Setup(q)

	// resale
	router.GET("/v1/app/users/:cpf/orders/:order_id/items", resaleContainer.Handler.GetOrderItemsByCPFAndOrderID)
	router.PUT("/v1/app/users/:cpf/orders/:order_id/items/:item_id/cancel", resaleContainer.Handler.CancelOrderItem)

	// retailer
	router.POST("/v1/retailer", retailerContainer.Handler.Create)
	router.GET("/v1/retailer", retailerContainer.Handler.List)
	router.GET("/v1/retailer/:id", retailerContainer.Handler.GetByID)
	router.PUT("/v1/retailer/:id", retailerContainer.Handler.Update)
	router.DELETE("/v1/retailer/:id", retailerContainer.Handler.Delete)
}

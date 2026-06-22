package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

type Handler struct {
	getOrderItemsUC   *usecase.GetOrderItemsUseCase
	cancelOrderItemUC *usecase.CancelOrderItemUseCase
}

func NewHandler(getUC *usecase.GetOrderItemsUseCase, cancelUC *usecase.CancelOrderItemUseCase) *Handler {
	return &Handler{
		getOrderItemsUC:   getUC,
		cancelOrderItemUC: cancelUC,
	}
}

func (h *Handler) GetOrderItemsByCPFAndOrderID(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")

	items, err := h.getOrderItemsUC.Execute(c.Request.Context(), cpf, orderID)
	if err != nil {
		if err.Error() == exception.ErrOrderNotFound.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *Handler) CancelOrderItem(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")
	itemID := c.Param("item_id")

	err := h.cancelOrderItemUC.Execute(c.Request.Context(), cpf, orderID, itemID)
	if err != nil {
		if err.Error() == exception.ErrOrderNotFound.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
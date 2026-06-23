package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	httpresponse "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/presentation/http/response"
)

type Handler struct {
	listUseCase   *usecase.ListResaleOrderItemsUseCase
	cancelUseCase *usecase.CancelResaleOrderItemUseCase
}

func NewHandler(listUseCase *usecase.ListResaleOrderItemsUseCase, cancelUseCase *usecase.CancelResaleOrderItemUseCase) *Handler {
	return &Handler{listUseCase: listUseCase, cancelUseCase: cancelUseCase}
}

// GetOrderItemsByCPFAndOrderID godoc
// @Summary List order items by CPF and order ID
// @Description Returns all items of a specific order identified by user CPF and order ID
// @Tags Resale
// @Accept json
// @Produce json
// @Param cpf path string true "User CPF (digits only, 11 characters)"
// @Param order_id path string true "Order ID (UUID)"
// @Success 200 {object} map[string]interface{} "Order items retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items [get]
func (h *Handler) GetOrderItemsByCPFAndOrderID(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")

	result, err := h.listUseCase.Execute(c.Request.Context(), orderID, cpf)
	if err != nil {
		if errors.Is(err, exception.ErrResaleOrderNotFound) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusOK, result.Items)
}

// CancelOrderItem godoc
// @Summary Cancel an order item
// @Description Cancels a specific item of an order identified by user CPF, order ID and item ID
// @Tags Resale
// @Accept json
// @Produce json
// @Param cpf path string true "User CPF (digits only, 11 characters)"
// @Param order_id path string true "Order ID (UUID)"
// @Param item_id path string true "Order Item ID (UUID)"
// @Success 200 {object} map[string]interface{} "Item cancelled successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order or item not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items/{item_id}/cancel [put]
func (h *Handler) CancelOrderItem(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")
	itemID := c.Param("item_id")

	err := h.cancelUseCase.Execute(c.Request.Context(), orderID, itemID, cpf)
	if err != nil {
		if errors.Is(err, exception.ErrCancelWindowExpired) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if errors.Is(err, exception.ErrResaleOrderItemNotFound) ||
			errors.Is(err, exception.ErrResaleOrderNotFound) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

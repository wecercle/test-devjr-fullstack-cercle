package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	httpresponse "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/presentation/http/response"
)

type Handler struct {
	listUseCase   *usecase.ListOrderItemsByCPFAndOrderIDUseCase
	cancelUseCase *usecase.CancelOrderItemUseCase
}

func NewHandler(listUseCase *usecase.ListOrderItemsByCPFAndOrderIDUseCase, cancelUseCase *usecase.CancelOrderItemUseCase) *Handler {
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
// @Success 200 {array} output.OrderItemOutputDTO "Order items retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items [get]
func (h *Handler) GetOrderItemsByCPFAndOrderID(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")

	result, err := h.listUseCase.Execute(c.Request.Context(), cpf, orderID)
	if err != nil {
		if isResaleBadRequestError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if isResaleNotFoundError(err) {
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
// @Success 204 "Item cancelled successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order or item not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items/{item_id}/cancel [put]
func (h *Handler) CancelOrderItem(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")
	itemID := c.Param("item_id")

	if err := h.cancelUseCase.Execute(c.Request.Context(), cpf, orderID, itemID); err != nil {
		if isResaleBadRequestError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if isResaleNotFoundError(err) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func isResaleBadRequestError(err error) bool {
	return errors.Is(err, exception.ErrInvalidOrderID) ||
		errors.Is(err, exception.ErrInvalidOrderItemID) ||
		errors.Is(err, exception.ErrOrderItemNotReturnable) ||
		errors.Is(err, exception.ErrOrderItemReturnExpired)
}

func isResaleNotFoundError(err error) bool {
	return errors.Is(err, exception.ErrOrderNotFound) ||
		errors.Is(err, exception.ErrOrderItemNotFound)
}

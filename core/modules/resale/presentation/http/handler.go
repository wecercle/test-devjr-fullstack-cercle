package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	httpresponse "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/presentation/http/response"
)

type Handler struct {
	getOrderItemsUseCase   *usecase.GetOrderItemsUseCase
	cancelOrderItemUseCase *usecase.CancelOrderItemUseCase
}

func NewHandler(getUC *usecase.GetOrderItemsUseCase, cancelUC *usecase.CancelOrderItemUseCase) *Handler {
	return &Handler{getOrderItemsUseCase: getUC, cancelOrderItemUseCase: cancelUC}
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
	cpf, orderID := c.Param("cpf"), c.Param("order_id")
	result, err := h.getOrderItemsUseCase.Execute(c.Request.Context(), cpf, orderID)
	if err != nil {
		h.handleGetError(c, err)
		return
	}
	httpresponse.Success(c, http.StatusOK, result)
}

func (h *Handler) handleGetError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domainException.ErrInvalidOrderID):
		httpresponse.DomainBadRequest(c, err)
	case errors.Is(err, domainException.ErrOrderNotFound):
		httpresponse.DomainNotFound(c, err)
	default:
		httpresponse.InternalServerError(c, err.Error())
	}
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
	cpf, orderID, itemID := c.Param("cpf"), c.Param("order_id"), c.Param("item_id")
	idempotent, err := h.cancelOrderItemUseCase.Execute(c.Request.Context(), cpf, orderID, itemID)
	if err != nil {
		h.handleCancelError(c, err)
		return
	}
	_ = idempotent
	c.Status(http.StatusNoContent)
}

func (h *Handler) handleCancelError(c *gin.Context, err error) {
	switch {
	case isValidationError(err):
		httpresponse.DomainBadRequest(c, err)
	case errors.Is(err, domainException.ErrOrderNotFound), errors.Is(err, domainException.ErrItemNotFound):
		httpresponse.DomainNotFound(c, err)
	default:
		httpresponse.InternalServerError(c, err.Error())
	}
}

func isValidationError(err error) bool {
	return errors.Is(err, domainException.ErrInvalidOrderID) ||
		errors.Is(err, domainException.ErrInvalidItemID) ||
		errors.Is(err, domainException.ErrItemNotEligibleForReturn)
}

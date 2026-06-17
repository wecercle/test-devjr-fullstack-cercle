package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/dto/input"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
	httpresponse "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/presentation/http/response"
)

type Handler struct {
	createUseCase  *usecase.CreateRetailerUseCase
	updateUseCase  *usecase.UpdateRetailerUseCase
	listUseCase    *usecase.ListRetailerUseCase
	getByIDUseCase *usecase.GetRetailerByIDUseCase
	deleteUseCase  *usecase.DeleteRetailerUseCase
}

func NewHandler(createUseCase *usecase.CreateRetailerUseCase, updateUseCase *usecase.UpdateRetailerUseCase, listUseCase *usecase.ListRetailerUseCase, getByIDUseCase *usecase.GetRetailerByIDUseCase, deleteUseCase *usecase.DeleteRetailerUseCase) *Handler {
	return &Handler{createUseCase: createUseCase, updateUseCase: updateUseCase, listUseCase: listUseCase, getByIDUseCase: getByIDUseCase, deleteUseCase: deleteUseCase}
}

// Create godoc
// @Summary Create a retailer
// @Description Creates a new retailer
// @Tags Retailer
// @Accept json
// @Produce json
// @Param body body input.CreateRetailerInputDTO true "Retailer data"
// @Success 201 {object} output.RetailerOutputDTO "Retailer created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/retailer [post]
func (h *Handler) Create(c *gin.Context) {
	var inputDTO input.CreateRetailerInputDTO
	if err := c.ShouldBindJSON(&inputDTO); err != nil {
		httpresponse.BadRequest(c, err.Error())
		return
	}

	result, err := h.createUseCase.Execute(c.Request.Context(), inputDTO)
	if err != nil {
		if isRetailerValidationError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if errors.Is(err, exception.ErrRetailerDocumentNumberExists) {
			httpresponse.DomainConflict(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusCreated, result)
}

// Update godoc
// @Summary Update a retailer
// @Description Updates an existing retailer by ID
// @Tags Retailer
// @Accept json
// @Produce json
// @Param id path string true "Retailer ID (UUID)"
// @Param body body input.UpdateRetailerInputDTO true "Updated retailer data"
// @Success 200 {object} output.RetailerOutputDTO "Retailer updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Retailer not found"
// @Failure 409 {object} map[string]string "Retailer is deleted"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/retailer/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var inputDTO input.UpdateRetailerInputDTO
	if err := c.ShouldBindJSON(&inputDTO); err != nil {
		httpresponse.BadRequest(c, err.Error())
		return
	}

	result, err := h.updateUseCase.Execute(c.Request.Context(), id, inputDTO)
	if err != nil {
		if isRetailerValidationError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if errors.Is(err, exception.ErrRetailerNotFound) || errors.Is(err, exception.ErrRetailerDeleted) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusOK, result)
}

// List godoc
// @Summary List retailers
// @Description Returns all active retailers
// @Tags Retailer
// @Produce json
// @Success 200 {array} output.RetailerOutputDTO "Retailers retrieved successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/retailer [get]
func (h *Handler) List(c *gin.Context) {
	result, err := h.listUseCase.Execute(c.Request.Context())
	if err != nil {
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusOK, result.Items)
}

// GetByID godoc
// @Summary Get a retailer by ID
// @Description Returns a specific retailer by ID
// @Tags Retailer
// @Produce json
// @Param id path string true "Retailer ID (UUID)"
// @Success 200 {object} output.RetailerOutputDTO "Retailer retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Retailer not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/retailer/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.getByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if isRetailerValidationError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if errors.Is(err, exception.ErrRetailerNotFound) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusOK, result)
}

// Delete godoc
// @Summary Delete a retailer
// @Description Soft deletes a retailer by ID
// @Tags Retailer
// @Produce json
// @Param id path string true "Retailer ID (UUID)"
// @Success 204 "Retailer deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Retailer not found"
// @Failure 409 {object} map[string]string "Retailer already deleted"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/retailer/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.deleteUseCase.Execute(c.Request.Context(), id); err != nil {
		if isRetailerValidationError(err) {
			httpresponse.DomainBadRequest(c, err)
			return
		}
		if errors.Is(err, exception.ErrRetailerNotFound) {
			httpresponse.DomainNotFound(c, err)
			return
		}
		if errors.Is(err, exception.ErrRetailerDeleted) {
			c.Status(http.StatusNoContent)
			return
		}
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func isRetailerValidationError(err error) bool {
	return errors.Is(err, exception.ErrInvalidRetailerID) ||
		errors.Is(err, exception.ErrInvalidRetailerDocumentNumber) ||
		errors.Is(err, exception.ErrInvalidRetailerName) ||
		errors.Is(err, exception.ErrInvalidRetailerTradeName)
}

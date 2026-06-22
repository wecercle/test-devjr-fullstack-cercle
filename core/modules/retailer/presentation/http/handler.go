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


func (h *Handler) List(c *gin.Context) {
	result, err := h.listUseCase.Execute(c.Request.Context())
	if err != nil {
		httpresponse.InternalServerError(c, err.Error())
		return
	}
	httpresponse.Success(c, http.StatusOK, result.Items)
}

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

package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	sharedexception "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/exception"
)

// DomainFailure é uma função auxiliar para responder a falhas usando o código de erro do domínio quando disponível.
// Se o erro não for um DomainError, ela responde com um código e título de fallback, e a mensagem de erro original.
func DomainFailure(c *gin.Context, statusCode int, fallbackCode, fallbackTitle string, err error) {
	var de *sharedexception.DomainError
	if errors.As(err, &de) {
		Failure(c, statusCode, de.Code, fallbackTitle, de.Message)
		return
	}
	Failure(c, statusCode, fallbackCode, fallbackTitle, err.Error())
}

// DomainBadRequest responde 400 usando o código de erro do domínio quando disponível.
func DomainBadRequest(c *gin.Context, err error) {
	DomainFailure(c, http.StatusBadRequest, "BAD_REQUEST", "Bad Request", err)
}

// DomainNotFound responde 404 usando o código de erro do domínio quando disponível.
func DomainNotFound(c *gin.Context, err error) {
	DomainFailure(c, http.StatusNotFound, "NOT_FOUND", "Not Found", err)
}

// DomainConflict responde 409 usando o código de erro do domínio quando disponível.
func DomainConflict(c *gin.Context, err error) {
	DomainFailure(c, http.StatusConflict, "CONFLICT", "Conflict", err)
}

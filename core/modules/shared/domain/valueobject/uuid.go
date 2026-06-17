package valueobject

import (
	"github.com/google/uuid"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/exception"
)

// UUID representa um identificador único validado como UUID v4, reutilizável em qualquer módulo.
type UUID struct {
	value string
}

// NewUUID cria um UUID value object a partir de uma string, validando que seja um UUID v4 válido.
func NewUUID(raw string) (UUID, error) {
	parsed, err := uuid.Parse(raw)
	if err != nil || parsed.Version() != 4 {
		return UUID{}, exception.ErrInvalidID
	}
	return UUID{value: parsed.String()}, nil
}

// String retorna a representação string do UUID.
func (u UUID) String() string { return u.value }

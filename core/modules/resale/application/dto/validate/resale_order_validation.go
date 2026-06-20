package validate

import (
	"github.com/google/uuid"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

func ValidateOrderID(id string) error {
	if id == "" {
		return domainException.ErrInvalidOrderID
	}
	if _, err := uuid.Parse(id); err != nil {
		return domainException.ErrInvalidOrderID
	}
	return nil
}

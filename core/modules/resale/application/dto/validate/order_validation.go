package validate

import (
	"github.com/google/uuid"
	"regexp"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

var cpfPattern = regexp.MustCompile(`^\d{11}$`)

func ValidateOrderID(id string) error {
	if id == "" {
		return domainException.ErrInvalidOrderID
	}
	if _, err := uuid.Parse(id); err != nil {
		return domainException.ErrInvalidOrderID
	}
	return nil
}

func ValidateItemID(id string) error {
	if id == "" {
		return domainException.ErrInvalidItemID
	}
	if _, err := uuid.Parse(id); err != nil {
		return domainException.ErrInvalidItemID
	}
	return nil
}

func ValidateCPF(cpf string) error {
	if !cpfPattern.MatchString(cpf) {
		return domainException.ErrInvalidCPF
	}
	return nil
}

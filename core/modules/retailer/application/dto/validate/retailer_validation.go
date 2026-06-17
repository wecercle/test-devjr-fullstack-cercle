package validate

import (
	"github.com/google/uuid"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
)

func ValidateRetailerID(id string) error {
	if id == "" {
		return domainException.ErrInvalidRetailerID
	}
	if _, err := uuid.Parse(id); err != nil {
		return domainException.ErrInvalidRetailerID
	}
	return nil
}

func ValidateRetailerRequiredFields(id, name, tradeName string) error {
	if err := ValidateRetailerID(id); err != nil {
		return err
	}
	if name == "" {
		return domainException.ErrInvalidRetailerName
	}
	if tradeName == "" {
		return domainException.ErrInvalidRetailerTradeName
	}
	return nil
}

// ValidateRetailerRequiredFieldsForCreate validates required fields for create (id is generated server-side).
// CNPJ format validation is enforced by the domain aggregate.
func ValidateRetailerRequiredFieldsForCreate(documentNumber, name, tradeName string) error {
	if documentNumber == "" {
		return domainException.ErrInvalidRetailerDocumentNumber
	}
	if name == "" {
		return domainException.ErrInvalidRetailerName
	}
	if tradeName == "" {
		return domainException.ErrInvalidRetailerTradeName
	}
	return nil
}

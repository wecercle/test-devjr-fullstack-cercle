package validate

import (
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

func ValidateUserDocumentNumber(documentNumber string) error {
	if documentNumber == "" {
		return domainException.ErrInvalidUserDocumentNumber
	}
	return nil
}

func ValidateOrderID(orderID string) error {
	if _, err := sharedValueobject.NewUUID(orderID); err != nil {
		return domainException.ErrInvalidOrderID
	}
	return nil
}

func ValidateOrderItemID(itemID string) error {
	if _, err := sharedValueobject.NewUUID(itemID); err != nil {
		return domainException.ErrInvalidOrderItemID
	}
	return nil
}

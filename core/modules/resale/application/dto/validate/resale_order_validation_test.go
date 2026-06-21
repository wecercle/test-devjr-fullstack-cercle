package validate

import (
	"errors"
	"testing"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
)

func TestValidateOrderID(t *testing.T) {
	if err := ValidateOrderID("6ae28fac-280e-4879-bbe3-9277decb4a06"); err != nil {
		t.Fatalf("expected valid order id, got %v", err)
	}

	if err := ValidateOrderID("invalid"); !errors.Is(err, domainException.ErrInvalidOrderID) {
		t.Fatalf("expected ErrInvalidOrderID, got %v", err)
	}
}

func TestValidateOrderItemID(t *testing.T) {
	if err := ValidateOrderItemID("b14e75f8-a3d6-4f8d-8a2b-5fbb0cf51001"); err != nil {
		t.Fatalf("expected valid order item id, got %v", err)
	}

	if err := ValidateOrderItemID("invalid"); !errors.Is(err, domainException.ErrInvalidOrderItemID) {
		t.Fatalf("expected ErrInvalidOrderItemID, got %v", err)
	}
}

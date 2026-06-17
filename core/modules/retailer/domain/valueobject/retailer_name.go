package valueobject

import (
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
)

const (
	retailerNameMinLen = 2
	retailerNameMaxLen = 100
)

// RetailerName representa o nome legal do varejista, validado como um value object. O nome deve ter entre 2 e 100 caracteres.
type RetailerName struct {
	value string
}

// NewRetailerName cria um RetailerName value object, aplicando as invariantes de comprimento.
func NewRetailerName(raw string) (RetailerName, error) {
	l := len([]rune(raw))
	if l < retailerNameMinLen || l > retailerNameMaxLen {
		return RetailerName{}, domainException.ErrInvalidRetailerName
	}
	return RetailerName{value: raw}, nil
}

// String retorna a representação string do nome do varejista.
func (n RetailerName) String() string { return n.value }

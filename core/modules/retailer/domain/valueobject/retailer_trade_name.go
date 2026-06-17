package valueobject

import (
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
)

const (
	retailerTradeNameMinLen = 2
	retailerTradeNameMaxLen = 25
)

// RetailerTradeName representa o nome fantasia do varejista, validado como um value object. O nome fantasia deve ter entre 2 e 25 caracteres.
type RetailerTradeName struct {
	value string
}

// NewRetailerTradeName cria um RetailerTradeName value object, aplicando as invariantes de comprimento.
func NewRetailerTradeName(raw string) (RetailerTradeName, error) {
	l := len([]rune(raw))
	if l < retailerTradeNameMinLen || l > retailerTradeNameMaxLen {
		return RetailerTradeName{}, domainException.ErrInvalidRetailerTradeName
	}
	return RetailerTradeName{value: raw}, nil
}

// String retorna a representação string do nome fantasia do varejista.
func (t RetailerTradeName) String() string { return t.value }

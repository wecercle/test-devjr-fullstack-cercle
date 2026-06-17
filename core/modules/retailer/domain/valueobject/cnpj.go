package valueobject

import (
	"regexp"
	"strconv"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
)

var nonDigit = regexp.MustCompile(`\D`)

// CNPJ representa o número do documento de identificação de pessoa jurídica no Brasil, validado e formatado como um value object.
// O CNPJ é composto por 14 dígitos, onde
// os 8 primeiros identificam a empresa,
// os 4 seguintes identificam a filial e
// os 2 últimos são dígitos verificadores calculados a partir dos 12 dígitos anteriores.
type CNPJ struct {
	value string
}

// NewCNPJ cria um CNPJ value object a partir de uma string bruta
func NewCNPJ(raw string) (CNPJ, error) {
	digits := nonDigit.ReplaceAllString(raw, "")
	if !isValidCNPJ(digits) {
		return CNPJ{}, domainException.ErrInvalidRetailerDocumentNumber
	}
	return CNPJ{value: digits}, nil
}

// String retorna a representação string do CNPJ, que é a sequência de 14 dígitos sem formatação.
func (c CNPJ) String() string { return c.value }

// isValidCNPJ valida se a string de dígitos representa um CNPJ válido
func isValidCNPJ(digits string) bool {
	if len(digits) != 14 {
		return false
	}

	// Rejeita números com todos os dígitos iguais (ex. "00000000000000")
	allSame := true
	for i := 1; i < 14; i++ {
		if digits[i] != digits[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	check1 := cnpjCheckDigit(digits[:12], weights1)
	check2 := cnpjCheckDigit(digits[:13], weights2)

	return digits[12] == byte('0'+check1) && digits[13] == byte('0'+check2)
}

// cnpjCheckDigit calcula o dígito verificador do CNPJ usando os pesos fornecidos
func cnpjCheckDigit(digits string, weights []int) int {
	sum := 0
	for i, w := range weights {
		d, _ := strconv.Atoi(string(digits[i]))
		sum += d * w
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

package cnpj

import (
	"strings"
)

const (
	_ = iota + 11
	CpnjBaseDigits
	CnpjCheckDigitInit
	CnpjSize

	// ASCII Value
	ASCII_Number = 48
)

var (
	weightsD1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weightsD2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsValid checks if the provided CNPJ is valid.
// It supports both the traditional numeric format and the new alphanumeric format.
func IsValid(cnpj string) bool {
	cnpj = sanitize(cnpj)

	if len(cnpj) != CnpjSize {
		return false
	}

	if isAllSame(cnpj) {
		return false
	}

	var c byte

	// First 12 characters can be alphanumeric.
	// Last 2 characters (check digits) must be numeric.
	for i := range CpnjBaseDigits {
		c = cnpj[i]
		if !(c >= '0' && c <= '9') && !(c >= 'A' && c <= 'Z') {
			return false
		}
	}

	for i := CpnjBaseDigits; i < CnpjSize; i++ {
		c = cnpj[i]
		if !(c >= '0' && c <= '9') {
			return false
		}
	}

	d1 := calculateDigit(cnpj[:CpnjBaseDigits], weightsD1)
	if d1 != int(cnpj[CpnjBaseDigits]-'0') {
		return false
	}

	d2 := calculateDigit(cnpj[:CnpjCheckDigitInit], weightsD2)
	return d2 == int(cnpj[CnpjCheckDigitInit]-'0')
}

func isAllSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// sanitize removes any formatting characters and converts to uppercase.
func sanitize(cnpj string) string {
	var sb strings.Builder
	for i := range len(cnpj) {
		c := cnpj[i]
		if c >= '0' && c <= '9' {
			sb.WriteByte(c)
		} else if c >= 'A' && c <= 'Z' {
			sb.WriteByte(c)
		} else if c >= 'a' && c <= 'z' {
			sb.WriteByte(c - 32)
		}
	}
	return sb.String()
}

// calculateDigit calculates a check digit based on the weights.
func calculateDigit(s string, weights []int) int {
	sum := 0
	for i := range len(s) {
		val := int(s[i]) - ASCII_Number
		sum += val * weights[i]
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

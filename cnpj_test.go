package cnpj

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		cnpj     string
		expected bool
	}{
		{"", false},
		{"12345678000195", true},
		{"12ABC34501AB77", true},
		{"AB12CD34EF5602", true},
		{"A1B2C3D4E5F668", true},
		{"ZXCVBN1234QW16", true},
		{"00000000000191", true},
		{"00000000000192", false},
		{"12345678000194", false},
		{"AB12CD34EF5601", false},
		{"INVALIDCNPJ123", false},
		{"12.345.678/0001-95", true},
		{"12abc34501ab77", true},
		{"00000000000000", false},
		{"11111111111111", false},
	}

	for _, tc := range testCases {
		t.Run(tc.cnpj, func(t *testing.T) {
			if got := IsValid(tc.cnpj); got != tc.expected {
				t.Errorf("IsValid(%q) = %v; want %v", tc.cnpj, got, tc.expected)
			}
		})
	}
}

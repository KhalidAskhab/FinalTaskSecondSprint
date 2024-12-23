package internal

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"1 + 1", 2, false},
		{"2 - 1", 1, false},
		{"2 * 3", 6, false},
		{"6 / 2", 3, false},
		{"2 + 3 * 4", 14, false},
		{"(1 + 2) * 3", 9, false},
		{"(1 + 2) * (3 + 4)", 21, false},
		{"(1 + (2 * 3) + 4)", 11, false},
		{"10 / (5 - 5)", 0, true}, // Деление на ноль
		{"(1 + 2", 0, true},       // Непарные скобки
		{"1 + 2)", 0, true},       // Непарные скобки
		{"1 + a", 0, true},        // Неправильный символ
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if (err != nil) != test.expectErr {
			t.Errorf("Calc(%q) expected error: %v, got: %v", test.expression, test.expectErr, err)
		}
		if !test.expectErr && result != test.expected {
			t.Errorf("Calc(%q) = %v; expected %v", test.expression, result, test.expected)
		}
	}
}

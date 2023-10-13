package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("CalculateTax(%f): expected %f, got %f", amount, expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	tests := []calcTax{
		{amount: 500, expected: 5.0},
		{amount: 1000, expected: 10.0},
		{amount: 1500, expected: 10.0},
	}

	for _, tt := range tests {
		result := CalculateTax(tt.amount)
		if result != tt.expected {
			t.Errorf("CalculateTax(%f): expected %f, got %f", tt.amount, tt.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTaxWithSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500.0)
	}
}

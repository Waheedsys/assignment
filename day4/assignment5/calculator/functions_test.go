package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}

	// Test negative result
	result = Subtract(4, 10)
	if result != -6 {
		t.Errorf("Expected -6, but got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 5)
	if result != 15 {
		t.Errorf("Expected 15, but got %d", result)
	}
}

func TestDivide(t *testing.T) {
	// Test valid division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Expected 5.0, but got %f", result)
	}

	// Test division by zero
	result, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error for divide by zero, but got nil")
	}
	if result != 0.0 {
		t.Errorf("Expected 0.0 for divide by zero, but got %f", result)
	}
}

func TestAddToLastValue(t *testing.T) {
	// First add
	addResult := Add(3, 5) // global = 8
	result := AddToLastValue(addResult, 2)
	if result != 10 {
		t.Errorf("Expected 10, but got %d", result)
	}

	// Add again
	result = AddToLastValue(result, 5)
	if result != 15 {
		t.Errorf("Expected 15, but got %d", result)
	}
}

func TestSubtractFromLastValue(t *testing.T) {
	// First subtract
	subtractResult := Subtract(10, 4) // global = 6
	result := SubtractFromLastValue(subtractResult, 3)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}

	// Subtract again
	result = SubtractFromLastValue(result, 2)
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

//benchmark

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3, 5)
	}
}

// BenchmarkSubtract benchmarks the Subtract function
func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(10, 4)
	}
}

// BenchmarkMultiply benchmarks the Multiply function
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(3, 5)
	}
}

// BenchmarkDivide benchmarks the Divide function
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(10, 2)
	}
}

// BenchmarkAddToLastValue benchmarks the AddToLastValue function
func BenchmarkAddToLastValue(b *testing.B) {
	lastResult := Add(3, 5) // Initial value set to 8
	for i := 0; i < b.N; i++ {
		AddToLastValue(lastResult, 2)
	}
}

// BenchmarkSubtractFromLastValue benchmarks the SubtractFromLastValue function
func BenchmarkSubtractFromLastValue(b *testing.B) {
	lastResult := Subtract(10, 4) // Initial value set to 6
	for i := 0; i < b.N; i++ {
		SubtractFromLastValue(lastResult, 3)
	}
}

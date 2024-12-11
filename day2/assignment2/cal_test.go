package assignment2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calculator(t *testing.T) {
	// test for "+"
	got, err := Calculator("+", 1.0, 2.1)
	assert.NoError(t, err)
	assert.Equal(t, 3.1, got, "Addition result mismatch")

	// test for "-"
	got, err = Calculator("-", 3.3, 1.4)
	assert.NoError(t, err)
	assert.Equal(t, 1.9, got, "subtraction result mismatch")

	// test for "*"
	got, err = Calculator("*", 2.6, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 6.5, got, "multiplication result mismatch")

	// test for "/"
	got, err = Calculator("/", 2.1, 2.8)
	assert.NoError(t, err)
	assert.InDelta(t, 0.75, got, 1e-9, "Division result mismatch")

	// test for "/"
	got, err = Calculator("/", 2.1, 0)
	assert.Error(t, err)
	assert.Equal(t, "cannot divide by zero", err.Error(), "Expected division by zero error")

	got, err = Calculator("o", 2, 3)
	assert.EqualError(t, err, "invalid operator o")
	assert.Equal(t, float64(0), got)
}

func BenchmarkCalculator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculator("+", 1.0, 2.1)
	}
}

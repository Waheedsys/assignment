package assignment2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	got, err := Sum(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, got)
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(3)
	}
}

package assignment1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Double(t *testing.T) {
	got := Double(2)
	assert.Equal(t, 4, got)
}

func BenchmarkDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Double(3)
	}
}

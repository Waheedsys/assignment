package assignment1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CirclePerimeter(t *testing.T) {
	got := CirclePerimeter(3)
	assert.Equal(t, 18.84, got)
}

func BenchmarkCirclePerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CirclePerimeter(3)
	}
}

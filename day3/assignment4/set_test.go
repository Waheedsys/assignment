package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSet(t *testing.T) {
	got := NewSet()
	assert.Equal(t, map[int]bool, got) // check
}

func BenchmarkNewSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSet()
	}
}

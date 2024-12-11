package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSet(t *testing.T) {
	got := NewSet()
	assert.NotNil(t, got)
}

func BenchmarkNewSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSet()
	}
}

func Test_AddIntoSet(t *testing.T) {
	got := AddIntoSet()
	assert.NotNil(t, got)
}

package assignment2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EvenOrodd(t *testing.T) {
	got := EvenOrodd(2)
	assert.Equal(t, "even", got)
	got = EvenOrodd(3)
	assert.Equal(t, "odd", got)

}

func BenchmarkEvenOrodd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenOrodd(2)
	}
}

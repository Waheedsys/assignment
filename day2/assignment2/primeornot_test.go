package assignment2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPrime(t *testing.T) {
	got := IsPrime(2)
	assert.Equal(t, true, got)
	got = IsPrime(4)
	assert.Equal(t, false, got)
	got = IsPrime(1)
	assert.Equal(t, false, got)

}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(2)
	}
}

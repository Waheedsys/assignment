package assignment5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stackt(t *testing.T) {
	got, err := Stack(8, []int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, got)
}

func BenchmarkStackt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Stack(8, []int{1, 2, 3})
	}
}

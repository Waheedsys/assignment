package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Reverse(t *testing.T) {
	got, err := Reverse([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 2, 1}, got)
}

func BenchmarkReverse(b *testing.B) {
	ans := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		Reverse(ans)
	}
}

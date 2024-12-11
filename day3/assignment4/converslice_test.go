package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Slicetomap(t *testing.T) {
	ans := []int{1, 2, 3, 4}
	a := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}
	got, err := Slicetomap(ans)
	assert.NoError(t, err)
	assert.Equal(t, a, got, "Something went wrong")
}

func BenchmarkSlicetomap(b *testing.B) {
	ans := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		Slicetomap(ans)
	}
}

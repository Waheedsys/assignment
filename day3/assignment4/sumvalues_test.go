package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SumValuesByKey(t *testing.T) {
	data := map[string][]int{
		"A": {1, 2, 3},
		"B": {4, 3, 2},
		"C": {8, 2, 6},
	}
	got, err := SumValuesByKey(data)
	assert.NoError(t, err, "somithing went wrong")
	expected := map[string]int{
		"A": 6,
		"B": 9,
		"C": 16,
	}
	assert.Equal(t, expected, got)
}

func BenchmarkSumValuesByKey(b *testing.B) {
	data := map[string][]int{
		"A": {1, 2, 3},
		"B": {4, 3, 2},
		"C": {8, 2, 6},
	}
	for i := 0; i < b.N; i++ {
		SumValuesByKey(data)
	}
}

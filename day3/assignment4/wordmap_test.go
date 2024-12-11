package assignment4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Takeword(t *testing.T) {
	got, err := Takeword("waheed")
	assert.NoError(t, err, "something went wrong")
	exepected := map[string]int{
		"a": 1,
		"d": 1,
		"e": 2,
		"h": 1,
		"w": 1,
	}
	assert.Equal(t, exepected, got)
}

func BenchmarkTakeword(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Takeword("bob")
	}
}

package assignment1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Greet(t *testing.T) {
	got := Greet("waheed")
	assert.Equal(t, "Hello,waheed", got)
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("bob")
	}
}

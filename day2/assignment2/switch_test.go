package assignment2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SwitchStatment(t *testing.T) {
	got := SwitchStatment(2)
	assert.Equal(t, "even", got)
	got = SwitchStatment(3)
	assert.Equal(t, "odd", got)
}

func BenchmarkSwitchStatment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwitchStatment(3)
	}
}

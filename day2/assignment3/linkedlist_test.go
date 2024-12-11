package assignment3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertAtBack(t *testing.T) {
	got := InsertAtBack(2)
	assert.Equal(t, "even", got)

}

func Test_DeleteLast(t *testing.T) {
	got := DeleteLast(2)
	assert.Equal(t, "even", got)

}

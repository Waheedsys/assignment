package assignment5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Push(t *testing.T) {
	stack := &Stack{}
	stack.Push(10)
	assert.Equal(t, 1, len(stack.GetStack()))
	assert.Equal(t, []int{10}, stack.GetStack())

	stack.Push(20)
	assert.Equal(t, 2, len(stack.GetStack()))
	assert.Equal(t, []int{10, 20}, stack.GetStack())
}

func BenchmarkPush(b *testing.B) {
	stack := &Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(12)
	}
}

func Test_Pop(t *testing.T) {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	got, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, got, "expected value 20")
	assert.Equal(t, 1, len(stack.GetStack()))
	assert.Equal(t, []int{10}, stack.GetStack())
}

func BenchmarkPop(b *testing.B) {
	stack := &Stack{}
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func TestGetStack(t *testing.T) {
	// Create a new stack
	stack := &Stack{}

	// Initially, the stack should be empty (nil or empty slice)
	actualStack := stack.GetStack()
	assert.True(t, len(actualStack) == 0 || actualStack == nil, "Expected stack to be empty initially")

	// Push some elements onto the stack
	stack.Push(10)
	stack.Push(20)

	// The stack should now contain the pushed elements
	assert.Equal(t, 2, len(stack.GetStack()))
	assert.Equal(t, []int{10, 20}, stack.GetStack())

	// Pop an element
	stack.Pop()

	// The stack should now contain only [10]
	assert.Equal(t, 1, len(stack.GetStack()))
	assert.Equal(t, []int{10}, stack.GetStack())
}

func BenchmarkGetStack(b *testing.B) {
	stack := &Stack{}
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {

		stack.GetStack()
	}
}

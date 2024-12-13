package assignment3

import (
	"fmt"
	"testing"
)

func TestInsertAtBack(t *testing.T) {
	list := &LinkedList{}
	list.InsertAtBack(1)
	list.InsertAtBack(2)
	list.InsertAtBack(3)

	expected := "1 -> 2 -> 3 -> "
	var actual string
	current := list.head
	for current != nil {
		actual += fmt.Sprintf("%d -> ", current.data)
		current = current.next
	}
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestDeleteLast(t *testing.T) {
	list := &LinkedList{}
	list.InsertAtBack(1)
	list.InsertAtBack(2)
	list.InsertAtBack(3)

	list.DeleteLast()

	expected := "1 -> 2 -> "
	var actual string
	current := list.head
	for current != nil {
		actual += fmt.Sprintf("%d -> ", current.data)
		current = current.next
	}
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDeleteLastEmptyList(t *testing.T) {
	list := &LinkedList{}
	list.DeleteLast()
	expected := ""
	var actual string
	current := list.head
	for current != nil {
		actual += fmt.Sprintf("%d -> ", current.data)
	}

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

// benchmark
func BenchmarkInsertAtBack(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < b.N; i++ {
		list.InsertAtBack(i)
	}
}

func BenchmarkDeleteLast(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < 1000; i++ {
		list.InsertAtBack(i)
	}
}

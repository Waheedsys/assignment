package assignment3

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}
func (list *LinkedList) DeleteLast() {
	// Check if the list is empty
	if list.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	// If there's only one node in the list
	if list.head.next == nil {
		list.head = nil
		fmt.Println("Last node of linked list has been deleted")
		return
	}

	// Traverse the list to find the second-to-last node
	current := list.head
	for current.next != nil && current.next.next != nil {
		current = current.next
	}

	// Remove the last node
	current.next = nil
	fmt.Println("Last node of linked list has been deleted")
}

func (list *LinkedList) Print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

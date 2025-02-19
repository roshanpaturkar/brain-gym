// Find the middle node of a linked list.
// Input: 0 1 2 3 4
// Output: 2
// If the number of nodes is even, then return the second middle node.
// Input: 0 1 2 3 4 5
// Output: 3

package main

import (
	"fmt"
)

type Node struct {
    val  int
    next *Node
}

func middleOfLinkedList(head *Node) *Node {
	slow := head
	fast := head
	
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	
	return slow
}

func main() {
	head := &Node{val: 1}
	head.next = &Node{val: 2}
	head.next.next = &Node{val: 3}
	head.next.next.next = &Node{val: 4}
	head.next.next.next.next = &Node{val: 5}
	
	fmt.Println(middleOfLinkedList(head).val)
	
	head.next.next.next.next.next = &Node{val: 6}
	fmt.Println(middleOfLinkedList(head).val)
	
	head.next.next.next.next.next.next = &Node{val: 7}
	fmt.Println(middleOfLinkedList(head).val)
}
package main

import "fmt"

type Node struct {
    val  int
    next *Node
}

func hasCycle(nodes *Node) bool {
	if nodes == nil || nodes.next == nil {
		return false
	}
	// Using Floyd's Cycle-Finding Algorithm (Tortoise and Hare)
	tortoise, hare := nodes, nodes.next

	for tortoise != hare {
		if hare == nil || hare.next == nil {
			return false
		}
		tortoise = tortoise.next
		hare = hare.next.next
	}

	return true
}

func main() {
	 // Example usage
	 node1 := &Node{val: 1}
	 node2 := &Node{val: 2}
	 node3 := &Node{val: 3}
	 node4 := &Node{val: 4}

	 // Creating a cycle
	 node1.next = node2
	 node2.next = node3
	 node3.next = node4
	 node4.next = node2 // Cycle here

	 fmt.Println(hasCycle(node1)) // Output: true
 
	 // Creating a non-cyclic list
	 node5 := &Node{val: 5}
	 node6 := &Node{val: 6}
	 node5.next = node6
 
	 fmt.Println(hasCycle(node5)) // Output: false
}
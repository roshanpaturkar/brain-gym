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

    slow, fast := nodes, nodes.next

    for slow != fast {
        if fast == nil || fast.next == nil {
            return false
        }

        slow = slow.next
        fast = fast.next.next
    }

    return true
}

func main() {
    node1 := &Node{val: 1}
    node2 := &Node{val: 2}
    node3 := &Node{val: 3}
    node4 := &Node{val: 4}

    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node2

    fmt.Println(hasCycle(node1))    // true

    // Creating a non-cyclic list
    node5 := &Node{val: 5}
    node6 := &Node{val: 6}
    node5.next = node6

    fmt.Println(hasCycle(node5))    // false
}
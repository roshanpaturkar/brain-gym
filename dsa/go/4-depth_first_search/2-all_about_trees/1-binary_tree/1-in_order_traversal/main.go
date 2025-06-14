package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
    val   int
    left  *Node
    right *Node
}


func inOrderTraversal(root *Node) {
    if root != nil {
        inOrderTraversal(root.left)
        fmt.Println(root.val)
        inOrderTraversal(root.right)
    }
}

func buildTree(nodes []string, pos *int) *Node {
    val := nodes[*pos]
    *pos++
    if val == "x" {
        return nil
    }
    v, _ := strconv.Atoi(val)
    left := buildTree(nodes, pos)
    right := buildTree(nodes, pos)
    return &Node{v, left, right}
}

func splitWords(s string) []string {
    if s == "" {
        return []string{}
    }
    return strings.Split(s, " ")
}


func main() {
	input := "5 4 3 x x 8 x x 6 x x"

	pos := 0
	root := buildTree(splitWords(input), &pos)

	fmt.Println("In-order traversal of the binary tree:")
	inOrderTraversal(root)

	input = "1 2 4 x x 5 x x 3 6 x x 7 x x"

	pos = 0
	root = buildTree(splitWords(input), &pos)
	fmt.Println("In-order traversal of the binary tree:")
	inOrderTraversal(root)
}
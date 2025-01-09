package main

import (
	"fmt"
)

func transpose(matrix [][]int) [][]int {
	var trans = [][]int{}
	for i := 0; i < len(matrix[0]); i++ {
		var row []int
		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[j][i])
		}
		trans = append(trans, row)
	}
	return trans
}


func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	trans := transpose(mat)
	fmt.Println(trans)

	mat = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(transpose(mat))

	mat = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(transpose(mat))
}

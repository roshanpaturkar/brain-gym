package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
    pos := []int{0, 0}
    for _, move := range moves {
        switch move {
            case 'R': pos[0]++
            case 'L': pos[0]--
            case 'U': pos[1]++
            case 'D': pos[1]--
        }
    }
    return (pos[0] == 0) && (pos[1] == 0)
}

func main() {
	robo := judgeCircle("UD")
	fmt.Println("Robo Returened: ", robo)

	robo = judgeCircle("LL")
	fmt.Println("Robo Returened: ", robo)

	robo = judgeCircle("UDLLURRD")
	fmt.Println("Robo Returened: ", robo)
}
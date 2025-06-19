package main

import "fmt"

func moveRobot(moves string) bool {
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
	robo := moveRobot("UD")
	fmt.Println("Robo Returened: ", robo)

	robo = moveRobot("LL")
	fmt.Println("Robo Returened: ", robo)

	robo = moveRobot("UDLLURRD")
	fmt.Println("Robo Returened: ", robo)
}
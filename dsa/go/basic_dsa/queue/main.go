package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func execute(program []string) []int {
    // initialize a slice representing a queue
    queue := []int{}
    for _, instruction := range program {
        switch instruction {
        case "peek":
            // print out the first item in the queue if it is not empty
            // print warning message if the queue is empty
            if len(queue) > 0 {
                fmt.Println(queue[0])
            } else {
                fmt.Println("Queue is empty!")
            }
        case "pop":
            // check if the queue is empty
            if len(queue) > 0 {
                // pop the first item in the queue
                queue = queue[1:]
            } else {
                // print message if the queue is empty
                fmt.Println("Queue is empty!")
            }
        default:
            // get the data in the "push" instruction
            data, err := strconv.Atoi(instruction[5:])
            if err == nil {
                // push data to the end of the queue
                queue = append(queue, data)
            }
        }
    }
    return queue
}

func arrayItoa(arr []int) []string {
    res := []string{}
    for _, v := range arr {
        res = append(res, strconv.Itoa(v))
    }
    return res
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    programLength, _ := strconv.Atoi(scanner.Text())
    program := []string{}
    for i := 0; i < programLength; i++ {
        scanner.Scan()
        program = append(program, scanner.Text())
    }
    res := execute(program)
    fmt.Println(strings.Join(arrayItoa(res), " "))
}

//8
// push 3
// push 7
// push 20
// peek
// pop
// push 0
// push 4
// pop
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func execute(program []string) []int {
    // initialize the stack
    var stack []int

    for _, instruction := range program {
        switch instruction {
        case "peek":
            // print out the top item in the stack
            if len(stack) > 0 {
                fmt.Println(stack[len(stack)-1])
            }
        case "pop":
            // pop the top item in the stack
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        default:
            // get the data in the "push" instruction
            data, err := strconv.Atoi(instruction[5:])
            if err == nil {
                // Push data to the top of the stack
                stack = append(stack, data)
            }
        }
    }
    return stack
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

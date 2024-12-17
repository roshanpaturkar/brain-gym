package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
    var  arr []string
    for i:=1; i<=n; i++ {
        if i%15 == 0 {
            arr = append(arr, "FizzBuzz")
        } else if i%3 == 0 && i%5 != 0 {
            arr = append(arr, "Fizz")
        } else if i%5 == 0 && i%3 != 0 {
            arr = append(arr, "Buzz")
        } else {
            arr = append(arr, strconv.Itoa(i))
        }
    }

    return arr
}

func main() {
	res := fizzBuzz(3)
	fmt.Println("N:3 =", res)

	res = fizzBuzz(5)
	fmt.Println("N:5 =", res)

	res = fizzBuzz(15)
	fmt.Println("N:15 =", res)
}
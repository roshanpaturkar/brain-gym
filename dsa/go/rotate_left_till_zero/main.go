package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func rotateLeftTillZero(nums []int) []int {
    for nums[0] != 0 {
        // remove front of queue and add it to end
        nums = append(nums[1:], nums[0])
    }
    return nums
}

func arrayAtoi(arr []string) []int {
    res := []int{}
    for _, x := range arr {
        v, _ := strconv.Atoi(x)
        res = append(res, v)
    }
    return res
}

func splitWords(s string) []string {
    if s == "" {
        return []string{}
    }
    return strings.Split(s, " ")
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
    nums := arrayAtoi(splitWords(scanner.Text()))
    res := rotateLeftTillZero(nums)
    fmt.Println(strings.Join(arrayItoa(res), " "))
}
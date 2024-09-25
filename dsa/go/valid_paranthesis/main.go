package main

import (
	"fmt"
	"strings"
)

func valid_parentheses(s string) bool {
	brackets := strings.Split(s, "")
	var brackets_stack []string

	for _, bracket := range brackets {
		if len(brackets_stack) != 0 {
			top := brackets_stack[len(brackets_stack)-1]

			if top == "(" && bracket == ")" {
				brackets_stack = brackets_stack[:len(brackets_stack)-1]
				continue
			}

			if top == "[" && bracket == "]" {
				brackets_stack = brackets_stack[:len(brackets_stack)-1]
				continue
			}

			if top == "{" && bracket == "}" {
				brackets_stack = brackets_stack[:len(brackets_stack)-1]
				continue
			}
		}
        brackets_stack = append(brackets_stack, bracket)
	}

	return len(brackets_stack) == 0 
}

func main() {
	valid := valid_parentheses("([])")
	fmt.Println(valid)

	valid = valid_parentheses("(]")
	fmt.Println(valid)

	valid = valid_parentheses("()[]{}")
	fmt.Println(valid)

	valid = valid_parentheses("()")
	fmt.Println(valid)
}
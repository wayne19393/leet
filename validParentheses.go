package main

import (
	"fmt"
)

func validParenthesis(p string) bool {
	stack := []rune{}
	for _, value := range p {
		switch value {
		case '(', '{', '[':
			stack = append(stack, value)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	tests := []string{
		"()[]{}", // true
		"([{}])", // true
		"(]",     // false
		"([)]",   // false
		"{[]}",   // true
		"",       // true
		"((((",   // false
		"(()())", // true
	}

	for _, t := range tests {
		fmt.Printf("%q => %v\n", t, validParenthesis(t))
	}
}

package main

import (
	_ "fmt"
)

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	start := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return start - i
}

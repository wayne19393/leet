package main

import (
	_ "fmt"
)

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		curr := roman[s[i]]

		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}
	return total
}

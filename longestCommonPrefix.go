package main

import (
	_ "fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	base := strs[0]

	for i := 0; i < len(base); i++ {
		for _, value := range strs[1:] {
			if i >= len(value) || value[i] != base[i] {
				return base[:i]
			}
		}
	}
	return base
}

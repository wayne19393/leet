package main

import (
	_ "fmt"
)

func removeDuplicateFromSortedArray(nums []int) int {
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

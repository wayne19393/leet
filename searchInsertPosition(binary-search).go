package main

import (
	_ "fmt"
)

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		k := i + (j-i)/2

		if nums[k] == target {
			return k
		} else if nums[k] < target {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return i
}

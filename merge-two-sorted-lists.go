package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Important part in which we implement the base logic of merging these two lists
func mergeToLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

// helper function to create lists
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}

func main() {
	// Example: list1 = 1 -> 3 -> 5
	list1 := createList([]int{1, 3, 5})
	// list2 = 2 -> 3 -> 4
	list2 := createList([]int{2, 3, 4})

	merged := mergeToLists(list1, list2)
	printList(merged) // Expected: 1 -> 2 -> 3 -> 3 -> 4 -> 5 -> nil
}

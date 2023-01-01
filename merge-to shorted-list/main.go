package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// create a dummy node to hold the result
	dummy := &ListNode{}

	// create a pointer to the result, which we will update as we go
	result := dummy

	// loop through both lists, adding the smaller element each time
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Next = list1
			list1 = list1.Next
		} else {
			result.Next = list2
			list2 = list2.Next
		}
		result = result.Next
	}

	// add the remaining elements from list1 or list2 (if any)
	if list1 != nil {
		result.Next = list1
	} else if list2 != nil {
		result.Next = list2
	}

	// return the next element of the dummy node, which is the head of the merged list
	return dummy.Next
}

func main() {
	// create two lists
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	// merge the lists
	merged := mergeTwoLists(list1, list2)

	// print the merged list
	for merged != nil {
		fmt.Println(merged.Val)
		merged = merged.Next
	}
}

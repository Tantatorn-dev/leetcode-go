package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConstructList(items []int) *ListNode {
	var head *ListNode
	var prev *ListNode

	for _, item := range items {
		node := new(ListNode)
		node.Val = item

		if head == nil {
			head = node
		} else {
			prev.Next = node
		}

		prev = node
	}

	return head
}

func PrintList(list *ListNode) {
	arr := []int{}

	for list != nil {
		arr = append(arr, list.Val)

		list = list.Next
	}

	fmt.Println(arr)
}

func CountList(list *ListNode) int {
	i := 0

	for list != nil {
		i += 1
		list = list.Next
	}

	return i
}
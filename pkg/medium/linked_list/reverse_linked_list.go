package linkedlist

import "leetcode-go/pkg/common"

func reverseBetween(head *common.ListNode,
	left int, right int) *common.ListNode {
	node := head

	var nums []int
	i := 0

	// first scan
	for node != nil {
		if i >= left-1 && i < right {
			nums = append(nums, node.Val)
		}

		node = node.Next
		i++
	}

	// second scan
	node = head
	i = 0
	for node != nil {
		if i >= left-1 && i < right {
			node.Val = nums[right-i-1]
		}

		node = node.Next
		i++
	}

	return head
}

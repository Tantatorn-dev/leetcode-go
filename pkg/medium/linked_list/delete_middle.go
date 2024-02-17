package linkedlist

import "leetcode-go/pkg/common"

func deleteMiddle(head *common.ListNode) *common.ListNode {
	node := head

	// find the length
	length := 0
	for node != nil {
		node = node.Next
		length++
	}

	if length == 1 {
		return nil
	}

	mid := length / 2

	node = head

	// let's remove mid element
	i := 0
	for node != nil {
		if i == mid-1 {
			node.Next = node.Next.Next
			break
		}

		i++
		node = node.Next
	}

	return head
}

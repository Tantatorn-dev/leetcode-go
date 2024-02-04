package medium

import "leetcode-go/pkg/common"

func reverseBetween(head *common.ListNode,
	left int, right int) *common.ListNode {
	node := head

	var leftLink *common.ListNode
	var rightLink *common.ListNode

	l := 0
	r := 0

	for node != nil {
		l++
		r++
		if l+1 == left {
			leftLink = node
		}
		if r+1 == right {
			rightLink = node.Next
		}

		node = node.Next
	}

	//temp := leftLink
	//leftLink.Next = rightLink
	//rightLink.Next = leftLink

	return head
}

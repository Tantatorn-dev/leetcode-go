package medium

import "leetcode-go/pkg/common"

func RemoveNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	length := common.CountList(head)

	targetIndex := length - n

	start := head

	i := 0

	var prev *common.ListNode

	for head != nil {
		if i == targetIndex {
			if prev != nil {
				prev.Next = head.Next
				break
			} else {
				if start.Next == nil {
					start = nil
				} else {
					start = head.Next
				}
				break
			}
		} else {
			prev = head
			head = head.Next
			i = i + 1
		}
	}

	return start
}

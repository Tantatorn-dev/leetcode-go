package medium

import "leetcode-go/pkg/common"

func detectCycle(head *common.ListNode) *common.ListNode {
	lsMap := make(map[*common.ListNode]*common.ListNode)

	for head != nil {
		if lsMap[head] == head {
			return head
		}

		lsMap[head] = head
		head = head.Next
	}

	return nil
}

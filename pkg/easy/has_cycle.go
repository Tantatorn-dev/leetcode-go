package easy

import "leetcode-go/pkg/common"

func hasCycle(head *common.ListNode) bool {
	lsMap := make(map[*common.ListNode]*common.ListNode)

	for head != nil {
		if lsMap[head] == head {
			return true
		}

		lsMap[head] = head
		head = head.Next
	}

	return false
}

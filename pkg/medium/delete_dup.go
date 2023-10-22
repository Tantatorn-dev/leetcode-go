package medium

import "leetcode-go/pkg/common"

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	dummy := new(common.ListNode)
	dummy.Next = head
	dummy.Val = -101

	lead := head
	follow := dummy

	for follow != nil && lead != nil {
		if lead.Next != nil && lead.Val == lead.Next.Val {
			temp := lead.Val

			for lead != nil && lead.Val == temp {
				lead = lead.Next
			}

			follow.Next = lead
		} else {
			lead = lead.Next
			follow = follow.Next
		}
	}

	return dummy.Next
}

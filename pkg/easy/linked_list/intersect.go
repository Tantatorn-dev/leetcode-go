package linkedlist

import "leetcode-go/pkg/common"

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	var arr []*common.ListNode

	for headA != nil {
		arr = append(arr, headA)
		headA = headA.Next
	}

	for headB != nil {
		for _, i := range arr {
			if i == headB {
				return headB
			}
		}

		headB = headB.Next
	}

	return nil
}

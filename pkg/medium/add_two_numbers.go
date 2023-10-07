package medium

import "leetcode-go/pkg/common"

func AddTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	carr := 0
	l1Start := l1

	if common.CountList(l1) < common.CountList(l2) {
		// swap l1 and l2
		temp := *l1
		*l1 = *l2
		*l2 = temp
	}

	for l1 != nil {
		if l1 != nil && l2 != nil {
			l1.Val = l1.Val + l2.Val + carr
		} else if l1 != nil {
			l1.Val = l1.Val + carr
		}

		if carr == 1 {
			carr = 0
		}

		if l1.Val > 9 {
			carr = 1
			l1.Val = l1.Val - 10
		}

		prevL1 := l1

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && carr > 0 {
			newNode := new(common.ListNode)
			newNode.Val = 1
			prevL1.Next = newNode
		}
	}

	return l1Start
}

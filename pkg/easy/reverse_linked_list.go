package easy

import "leetcode-go/pkg/common"

func reverseList(head *common.ListNode) *common.ListNode {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)

		head = head.Next
	}

	var ret *common.ListNode
	var prev *common.ListNode

	for i := len(arr) - 1; i >= 0; i-- {
		node := new(common.ListNode)
		node.Val = arr[i]

		if ret == nil {
			ret = node
		} else {
			prev.Next = node
		}

		prev = node
	}

	return ret
}

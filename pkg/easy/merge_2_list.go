package easy

import "leetcode-go/pkg/common"

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	newList := new(common.ListNode)
	dummyHead := newList

	for list1 != nil || list2 != nil {
		if list1 == nil {
			newList.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			newList.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			newList.Next = list2
			list2 = list2.Next
		} else {
			newList.Next = list1
			list1 = list1.Next
		}
		newList = newList.Next
	}

	return dummyHead.Next
}

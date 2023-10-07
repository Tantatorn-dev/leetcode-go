package main

import (
	"leetcode-go/pkg/common"
	"leetcode-go/pkg/medium"
)

func main() {
	myList1 := common.ConstructList([]int{1, 2, 3, 4})
	myList2 := common.ConstructList([]int{1, 2, 4, 5, 5, 3})

	sum := medium.AddTwoNumbers(myList1, myList2)

	common.PrintList(sum)
}

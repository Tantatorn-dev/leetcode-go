package easy

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	var ret []string

	var arr []int

	for i := 0; i < len(nums); i++ {
		var currentDiff int
		if i == len(nums)-1 {
			currentDiff = 0
		} else {
			currentDiff = nums[i+1] - nums[i]
		}
		arr = append(arr, nums[i])

		if currentDiff != 1 {
			if len(arr) > 1 {
				ret = append(ret, fmt.Sprintf("%d->%d", arr[0], arr[len(arr)-1]))
			} else {
				ret = append(ret, fmt.Sprintf("%d", arr[0]))
			}
			arr = []int{}
		}
	}

	return ret
}

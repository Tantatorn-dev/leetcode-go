package easy

// searchInsert: Binary Search
func searchInsert(nums []int, target int) int {
	mid := 0
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid = (high + low) / 2

		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	if mid > high {
		return mid
	}

	return mid + 1
}

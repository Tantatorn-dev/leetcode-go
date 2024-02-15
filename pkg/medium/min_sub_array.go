package medium

func minSubArrayLen(target int, nums []int) int {
	for windowSize := 1; windowSize < len(nums); windowSize++ {
		for i := 0; i < len(nums)-windowSize+1; i++ {
			if currSum := sum(nums[i : i+windowSize]); target == currSum {
				return windowSize
			}
		}
	}

	return 0
}

func sum(nums []int) int {
	ret := 0

	for _, n := range nums {
		ret += n
	}

	return ret
}

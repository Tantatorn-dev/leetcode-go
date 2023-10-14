package medium

func rotate(nums []int, k int) {
	l := len(nums)

	k = k % l

	nums = append(nums[l-k:], nums[:l-k]...)
}

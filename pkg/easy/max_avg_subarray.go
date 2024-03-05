package easy

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, n := range nums[:k] {
		sum += n
	}

	maxSum := sum

	for i := 1; i <= len(nums)-k; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

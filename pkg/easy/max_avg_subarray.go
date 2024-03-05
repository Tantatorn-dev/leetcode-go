package easy

func findMaxAverage(nums []int, k int) float64 {
	var maxSum *int

	for i := 0; i <= len(nums)-k; i++ {
		sum := 0

		for j := i; j < i+k; j++ {
			sum += nums[j]
		}

		if maxSum == nil || sum > *maxSum {
			maxSum = &sum
		}
	}

	return float64(*maxSum) / float64(k)
}

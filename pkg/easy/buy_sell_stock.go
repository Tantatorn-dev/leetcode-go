package easy

// O(N**2) solution
func maxProfit(prices []int) int {
	maxDiff := 0

	for i := 0; i < len(prices); i++ {
		// scan through remaining
		for j := i + 1; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	return maxDiff
}

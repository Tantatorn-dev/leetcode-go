package dp

func climbStairs(n int) int {
	dp := []int{1, 1}

	for i := 2; i < n+1; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[n]
}

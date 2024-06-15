package easy

func largestAltitude(gain []int) int {
	max := 0
	curr := 0

	for _, g := range gain {
		curr += g
		if curr > max {
			max = curr
		}
	}

	return max
}

package medium

func maxArea(height []int) int {
	max := 0

	i := 0
	j := len(height) - 1

	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > max {
			max = area
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return max
}

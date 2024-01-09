package medium

// take too much time; need to use bit shifting
func divide(dividend int, divisor int) int {
	count := 0

	if dividend < 0 && divisor < 0 {
		for remainder := dividend; remainder <= divisor; remainder -= divisor {
			count++
		}
	} else if divisor < 0 {
		for remainder := dividend; remainder >= -divisor; remainder += divisor {
			count--
		}
	} else if dividend < 0 {
		for remainder := dividend; remainder <= -divisor; remainder += divisor {
			count--
		}
	} else {
		for remainder := dividend; remainder >= divisor; remainder -= divisor {
			count++
		}
	}

	return count
}

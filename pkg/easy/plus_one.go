package easy

// Problem: https://leetcode.com/problems/plus-one
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] += 1
		}

		if digits[i] >= 10 {
			digits[i] = digits[i] - 10

			if i-1 >= 0 {
				digits[i-1] = digits[i-1] + 1
			} else {
				// prepend 1
				digits = append([]int{1}, digits...)
			}
		}
	}

	return digits
}

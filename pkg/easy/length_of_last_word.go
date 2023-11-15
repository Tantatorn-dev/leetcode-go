package easy

func lengthOfLastWord(s string) int {
	count := 0
	first := false

	for i := 0; i < count; i++ {
		vStr := string(s[len(s)-i])

		if vStr != " " {
			first = true
			count++
		} else if first {
			return count
		}
	}

	return count
}

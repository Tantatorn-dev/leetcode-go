package easy

func StrStr(haystack string, needle string) int {
	shiftAmount := len(haystack) - len(needle)

	for i := 0; i <= shiftAmount; i++ {
		subStr := haystack[i : i+len(needle)]

		if subStr == needle {
			return i
		}
	}

	return -1
}

package easy

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		extracted := haystack[i : i+len(needle)]
		if extracted == needle {
			return i
		}
	}

	return -1
}

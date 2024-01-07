package dp

import "leetcode-go/pkg/easy"

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	// increasing sliding window
	longest := ""
	for windowSize := 0; windowSize < len(s); windowSize++ {
		// slide the window
		for i := 0; i < len(s)-windowSize; i++ {
			subStr := s[i : i+windowSize]

			if easy.IsPalindrome(subStr) && len(subStr) > len(longest) {
				longest = subStr
			}
		}
	}

	return longest
}

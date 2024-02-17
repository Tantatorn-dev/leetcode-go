package easy

func reverseVowels(s string) string {
	var vowels []string
	var vowelIndexes []int

	// scan to get all vowels
	for i, ch := range s {
		chStr := string(ch)

		if isVowel(chStr) {
			vowels = append(vowels, chStr)
			vowelIndexes = append(vowelIndexes, i)
		}
	}

	// replace vowel
	for i := 0; i < len(vowels); i++ {
		index := vowelIndexes[i]
		s = s[:index] + vowels[len(vowels)-1-i] + s[index+1:]
	}

	return s
}

func isVowel(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u",
		"A", "E", "I", "O", "U"}

	for _, vw := range vowels {
		if s == vw {
			return true
		}
	}

	return false
}

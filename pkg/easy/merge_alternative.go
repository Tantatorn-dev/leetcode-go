package easy

func mergeAlternately(word1 string, word2 string) string {
	var length int
	if len(word1) > len(word2) {
		length = len(word1)
	} else {
		length = len(word2)
	}

	var ret string

	for i := 0; i < length; i++ {
		if i < len(word1) {
			ret += string(word1[i])
		}

		if i < len(word2) {
			ret += string(word2[i])
		}
	}

	return ret
}

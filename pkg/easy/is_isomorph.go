package easy

func isIsomorphic(s string, t string) bool {
	strMap1 := make(map[string]string)
	strMap2 := make(map[string]string)

	for i := 0; i < len(s); i++ {
		// if not exist before
		strMap1[string(s[i])] = string(t[i])
		strMap2[string(t[i])] = string(s[i])
	}

	newStr1 := ""
	newStr2 := ""
	for i := 0; i < len(t); i++ {
		newStr1 = newStr1 + strMap1[string(s[i])]
		newStr2 = newStr2 + strMap2[string(t[i])]
	}

	return newStr1 == t && newStr2 == s
}

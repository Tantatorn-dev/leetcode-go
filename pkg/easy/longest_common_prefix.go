package easy

func longestCommonPrefix(strs []string) string {
	ret := ""

outer:
	for i := 0; i < len(strs[0]); i++ {
		temp := strs[0][:i+1]

		for _, str := range strs[1:] {
			if len(str) > i && temp == str[:i+1] {
				continue
			} else {
				break outer
			}
		}

		ret = temp
	}

	return ret
}

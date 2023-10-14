package medium

func reverseWords(s string) string {
	ret := ""
	sArr := split(s)

	for i := 0; i < len(sArr); i++ {
		pre := ""
		if i > 0 {
			pre = " "
		}

		ret += pre + sArr[len(sArr)-i-1]
	}

	return ret
}

func split(s string) []string {
	ret := []string{"dummy"}

	temp := ""
	for _, ch := range s {
		chStr := string(ch)

		if chStr != " " {
			temp += chStr
		} else {
			if temp != "" && temp != ret[len(ret)-1] {
				ret = append(ret, temp)
				temp = ""
			}
		}
	}

	if temp != "" {
		ret = append(ret, temp)
	}

	return ret[1:]
}

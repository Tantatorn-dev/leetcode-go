package medium

var letterMap = map[string]string{
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// like in old phone
func LetterCombinations(digits string) []string {
	ret := []string{}

	for i, d := range digits {
		dStr := string(d)

		phoneKeys := letterMap[dStr]

		for _, pk := range phoneKeys {
			pkStr := string(pk)

			if i == 0 {
				ret = append(ret, pkStr)
			} else {
				for i := range ret {
					ret[i] += pkStr
				}
			}

		}
	}

	return ret
}

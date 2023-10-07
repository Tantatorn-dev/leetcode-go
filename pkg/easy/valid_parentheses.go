package easy

func IsValid(s string) bool {
	var stack []string

	for _, c := range s {
		cStr := string(c)

		if cStr == "(" || cStr == "[" || cStr == "{" {
			stack = append(stack, cStr)
		}

		if cStr == ")" || cStr == "]" || cStr == "}" {
			if len(stack) == 0 {
				return false
			} else {
				tail := stack[len(stack)-1]

				if tail == "(" && cStr == ")" {
					stack = stack[:len(stack)-1]
				} else if tail == "[" && cStr == "]" {
					stack = stack[:len(stack)-1]
				} else if tail == "{" && cStr == "}" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

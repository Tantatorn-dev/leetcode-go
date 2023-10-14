package easy

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		temp := a
		a = b
		b = temp
	}

	ret := ""
	carry := "0"

	for i := 0; i < len(a); i++ {
		if len(a) > i {
			aVal := string(a[len(a)-i-1])
			bVal := "0"

			if len(b) > i {
				bVal = string(b[len(b)-i-1])
			}

			aVal, carry = sumChar(aVal, bVal, carry)

			ret = aVal + ret
		}
	}

	if carry == "1" {
		ret = "1" + ret
	}

	return ret
}

func sumChar(aChar string, bChar string, c string) (res string, carry string) {
	if aChar == "1" && bChar == "1" {
		if c == "1" {
			return "1", "1"
		}
		return "0", "1"
	} else if aChar == "1" || bChar == "1" {
		if c == "1" {
			return "0", "1"
		}
		return "1", "0"
	} else {
		if c == "1" {
			return "1", "0"
		}
		return "0", "0"
	}
}

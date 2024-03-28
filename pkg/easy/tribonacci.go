package easy

func tribonacci(n int) int {
	a, b, c := 0, 1, 1

	if n == 0 {
		return 0
	}

	for i := 2; i < n; i++ {
		temp := a + b + c
		a = b
		b = c
		c = temp
	}

	return c
}

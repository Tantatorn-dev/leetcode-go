package easy

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}

	ret := make([]bool, len(candies))

	for i, candy := range candies {
		candy += extraCandies
		ret[i] = candy >= maxCandy
	}

	return ret
}

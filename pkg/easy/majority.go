package easy

func majorityElement(nums []int) int {
	mymap := make(map[int]int)

	for _, i := range nums {
		mymap[i] += 1
	}

	max := 0
	ret := 0
	for k, v := range mymap {
		if v >= max {
			max = v
			ret = k
		}
	}

	return ret
}

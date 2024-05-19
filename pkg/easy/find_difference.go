package easy

// https://leetcode.com/problems/find-the-difference-of-two-arrays
func findDifference(nums1 []int, nums2 []int) [][]int {
	numsMap1 := make(map[int]int)
	numsMap2 := make(map[int]int)

	for _, num := range nums1 {
		numsMap1[num]++
	}

	for _, num := range nums2 {
		numsMap2[num]++
	}

	ret := [][]int{{}, {}}

	for key := range numsMap1 {
		if _, ok := numsMap2[key]; !ok {
			ret[0] = append(ret[0], key)
		}
	}

	for key := range numsMap2 {
		if _, ok := numsMap1[key]; !ok {
			ret[1] = append(ret[1], key)
		}
	}

	return ret
}

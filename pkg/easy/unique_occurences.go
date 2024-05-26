package easy

// Problem: https://leetcode.com/problems/unique-number-of-occurrences
func uniqueOccurrences(arr []int) bool {
	numMap := make(map[int]int)

	for _, i := range arr {
		numMap[i]++
	}

	countMap := make(map[int]int)
	for _, val := range numMap {
		if _, ok := countMap[val]; ok {
			return false
		} else {
			countMap[val] = val
		}
	}

	return true
}

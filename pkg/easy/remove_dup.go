package easy

func removeDuplicates(nums []int) int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}

		numsMap[nums[i]]++
	}

	return len(nums)
}

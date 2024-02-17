package easy

func canPlaceFlowers(flowerbed []int, n int) bool {
	slot := 0

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}

	for i, flower := range flowerbed {
		if flower == 0 {
			if i == 0 && len(flowerbed) > i+1 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				slot++
			} else if i == len(flowerbed)-1 && i-1 > 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				slot++
			} else if i-1 > 0 && len(flowerbed) > i+1 &&
				flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				slot++
			}
		}
	}

	return slot >= n
}

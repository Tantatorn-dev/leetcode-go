package common

func QuickSort(arr []int) []int {
	// Base Case
	if len(arr) < 2 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			return arr
		}
	}

	pivot := arr[0]

	var lesser []int
	var greater []int

	for _, i := range arr[1:] {
		if i > pivot {
			greater = append(greater, i)
		} else {
			lesser = append(lesser, i)
		}
	}

	ret := append(QuickSort(lesser), pivot)
	return append(ret, QuickSort(greater)...)
}

func SelectionSort(arr []int) []int {
	var ret []int

	for i := 0; i < len(arr); i++ {

		min := arr[i]
		minIndex := i

		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		ret = append(ret, min)
	}

	return ret
}

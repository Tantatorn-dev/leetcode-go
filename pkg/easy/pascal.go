package easy

func generatePascal(numRows int) [][]int {
	var ret [][]int

	for i := 1; i <= numRows; i++ {
		arr := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				arr[j] = 1
			} else {
				prevArr := ret[i-2]
				arr[j] = prevArr[j-1] + prevArr[j]
			}
		}

		ret = append(ret, arr)
	}

	return ret
}

func getRow(rowIndex int) []int {
	var row []int

	for i := 1; i <= rowIndex+1; i++ {
		prev := row
		row = make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				row[j] = 1
			} else {
				row[j] = prev[j-1] + prev[j]
			}
		}
	}

	return row
}

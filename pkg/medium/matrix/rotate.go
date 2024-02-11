package matrix

func rotate(matrix [][]int) {
	length := len(matrix)

	temp := make([][]int, length)

	for i := 0; i < length; i++ {
		temp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			temp[i][j] = matrix[length-j-1][i]
		}
	}

	// rewrite temp to matrix
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = temp[i][j]
		}
	}
}

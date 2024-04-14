package matrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix[0]), len(matrix)

	low := 0
	high := m*n - 1

	for low <= high {
		mid := (high + low) / 2

		y := mid / m
		x := mid % m

		if matrix[y][x] == target {
			return true
		}

		if matrix[y][x] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

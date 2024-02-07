package matrix

func searchMatrix(matrix [][]int, target int) bool {
	ydim := len(matrix)
	xdim := len(matrix[0])

	midY := ydim / 2
	midX := xdim / 2

	for {
		val := matrix[midX][midY]

		if target == val {
			return true
		} else if target > val {
			if midX+1 < xdim {
				midX++
			} else if midY+1 < ydim {
				midY++
			} else {
				break
			}
		} else {
			if midX-1 >= 0 {
				midX--
			} else if midY-1 >= 0 {
				midY--
			} else {
				break
			}
		}
	}

	return false
}

package matrix

// https://leetcode.com/problems/set-matrix-zeroes/?envType=study-plan-v2&envId=top-interview-150
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var targets [][]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				targets = append(targets, []int{i, j})
			}
		}
	}

	for _, target := range targets {
		// set zeroes on horizontal
		for x := 0; x < n; x++ {
			matrix[target[0]][x] = 0
		}

		// set zeroes on vertical
		for x := 0; x < m; x++ {
			matrix[x][target[1]] = 0
		}
	}
}

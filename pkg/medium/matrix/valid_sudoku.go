package matrix

func isValidSudoku(board [][]byte) bool {
	return isRowValid(board) && isColumnValid(board) && isSubBoxesValid(board)
}

func isRowValid(board [][]byte) bool {
	for _, row := range board {
		seen := make(map[byte]byte)

		for _, i := range row {
			if i == '.' {
				continue
			}

			if _, ok := seen[i]; ok {
				return false
			} else {
				seen[i] = i
			}
		}
	}

	return true
}

func isColumnValid(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		seen := make(map[byte]byte)
		for j := 0; j < 9; j++ {
			val := board[j][i]

			if val == '.' {
				continue
			}

			if _, ok := seen[val]; ok {
				return false
			} else {
				seen[val] = val
			}
		}
	}

	return true
}

func isSubBoxesValid(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			seen := make(map[byte]byte)

			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					val := board[3*i+x][3*j+y]

					if val == '.' {
						continue
					}

					if _, ok := seen[val]; ok {
						return false
					} else {
						seen[val] = val
					}
				}
			}
		}
	}

	return true
}

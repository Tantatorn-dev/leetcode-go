package matrix

import (
	"fmt"
	"strings"
)

func equalPairs(grid [][]int) int {
	strMap := make(map[string]int)

	// scan row
	for _, row := range grid {
		key := makeKey(row)

		strMap[key]++
	}

	count := 0

	// scan col
	for i := 0; i < len(grid); i++ {
		var col []int
		for j := 0; j < len(grid); j++ {
			col = append(col, grid[j][i])
		}

		key := makeKey(col)

		if val, ok := strMap[key]; ok {
			count += val
		}
	}

	return count
}

func makeKey(input []int) string {
	var strs []string
	for _, i := range input {
		str := fmt.Sprintf("%d", i)
		strs = append(strs, str)
	}

	return strings.Join(strs, ",")
}

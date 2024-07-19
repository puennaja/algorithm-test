package diagonaldifference

// there have a metrix find the abs(diagonal left - diagonal right)

// 1 2 3
// 4 5 6
// 9 8 9

// diagolnal left-right = 1 + 5 + 9 = 15
// diagonal right-left = 3 + 5 + 9 = 17
// abs(15-17) == 2

func DiagonalDifference(arr [][]int32) int32 {
	if len(arr) == 0 {
		return 0
	}

	i := 0
	j := len(arr[0]) - 1

	var countLeft int32 = 0
	var countRight int32 = 0
	for _, row := range arr {
		countLeft += row[i]
		countRight += row[j]
		i++
		j--
	}

	// return abs
	if countLeft > countRight {
		return countLeft - countRight
	}
	return countRight - countLeft
}

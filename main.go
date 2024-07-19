package main

import (
	"algorithm-test/hackerrank/diagonaldifference"
)

func main() {
	// row 3 col 4
	// 00 03
	// 11 12
	// 22 21

	arr1 := []int32{1, 2, 3, 4}
	arr2 := []int32{1, 2, 3, 4}
	arr3 := []int32{1, 2, 3, 4}
	arr := [][]int32{arr1, arr2, arr3}
	diagonaldifference.DiagonalDifference(arr)
}

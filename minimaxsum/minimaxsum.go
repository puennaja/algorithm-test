package minimaxsum

import (
	"fmt"
	"slices"
)

// out print min sum and max sum of arr

func miniMaxSum(arr []int32) {
	// Write your code here
	slices.Sort(arr)
	sum := 0
	for _, v := range arr {
		sum += int(v)
	}
	fmt.Println(sum-int(arr[0]), sum-int(arr[len(arr)-1]))
}

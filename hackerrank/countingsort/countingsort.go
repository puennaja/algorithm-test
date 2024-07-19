package countingsort

/*


 */

func CountingSort(arr []int32) []int32 {
	frequencyArr := make([]int32, 100)
	for _, v := range arr {
		frequencyArr[v]++
	}

	sortedArr := []int32{}
	for i, v := range frequencyArr {
		if v == 0 {
			continue
		}
		for j := v; j != 0; j-- {
			sortedArr = append(sortedArr, int32(i))
		}
	}
	return sortedArr
}

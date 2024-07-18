package plusminus

import "fmt"

// arr = [1,1,1,0,-1,-1]
// plus => 3 in 6
// minus => 2 in 6
// zero => 1 in 6

// out
// print 6 decimal of plus / all
// print 6 decimal of minus / all
// print 6 decimal of zero / all

func PlusMinus(arr []int32) {
	lenght := len(arr)
	mapPlusMinus := map[string]int{
		"plus":  0,
		"minus": 0,
		"zero":  0,
	}
	for _, v := range arr {
		if v > 0 {
			mapPlusMinus["plus"]++
		} else if v == 0 {
			mapPlusMinus["zero"]++
		} else {
			mapPlusMinus["minus"]++
		}
	}

	fmt.Printf("%.6f\n", float64(mapPlusMinus["plus"])/float64(lenght))
	fmt.Printf("%.6f\n", float64(mapPlusMinus["minus"])/float64(lenght))
	fmt.Printf("%.6f\n", float64(mapPlusMinus["zero"])/float64(lenght))
}

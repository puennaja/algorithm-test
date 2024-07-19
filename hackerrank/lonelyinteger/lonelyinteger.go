package lonelyinteger

// find the unique integer in a arrays
// there is only one unique integer

// time: O(n) space: O(n)
func Lonelyinteger(a []int32) int32 {
	mapA := map[int32]int32{}
	for _, v := range a {
		mapA[v] += 1
	}

	for i, v := range mapA {
		if v < 2 {
			return i
		}
	}
	return -1
}

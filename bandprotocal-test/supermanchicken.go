package bandprotocaltest

// there have a input parameter
// chicken => number of the chickens
// supermanRoof => denites the length of roof that superman can rescue
// chickenPosition => position of the chicken in one-dimensional axis. (1 checken/position)

// output => maxinum number of chickens that superman can protect with the given roof length.

// note
/*
- superman can start at any point of roof
- the roof cover the chicken that are within supermanRoof [p, p+supermanRoof)
  if p = 1 and supermanRoof = 5 => roof can cover position 1,2,3,4,5
- the position of thet chicken will be sorted from lowest to highest
*/

// 5 5
// 2 5 10 12 15
// out 2
func SupermanChickenRescue(chicken, supermanRoof int, chickenPosition []int) int {
	maxChickenRescue := 0
	for i, position := range chickenPosition {
		roofRange := position + supermanRoof
		maxRescueRound := 0
		for _, p := range chickenPosition[i:] {
			if p < roofRange {
				maxRescueRound++
			} else {
				break
			}
		}

		if maxRescueRound > maxChickenRescue {
			maxChickenRescue = maxRescueRound
		}

		if maxChickenRescue == supermanRoof {
			break
		}
	}
	return maxChickenRescue
}

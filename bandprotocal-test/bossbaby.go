package bandprotocaltest

import "strings"

// input s with => "SSRRRSRS"
// S represent sort
// R represent Revenge

// output
// any S should have be R after once => Good Boy
// R can not be first => Bad Boy

const (
	GoodBoy string = "Good Boy"
	BadBoy  string = "Bad Boy"
)

func BossBabyRevenge(s string) string {
	strs := strings.Split(strings.ToUpper(s), "")
	count := 0
	for i, v := range strs {
		if i == 0 && v == "R" {
			count++
			break
		}
		if v == "S" {
			count++
		} else {
			if count != 0 {
				count--
			}
		}
	}

	if count <= 0 {
		return GoodBoy
	}
	return BadBoy

}

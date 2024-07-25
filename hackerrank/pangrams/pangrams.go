package pangrams

import (
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

// s is in a-z, A-Z and space
// check it is pangrams or not (pangram is have all english alphabet)
// yes return pangram
// no return not pangram

func Pangrams(s string) string {
	mapAlphabet := make(map[string]int32, 26)
	countAlphabet := int32(26)
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	strs := strings.Split(s, "")
	for _, v := range strs {
		if mapAlphabet[v] == int32(0) {
			mapAlphabet[v]++
			countAlphabet--
		}
		if countAlphabet <= int32(0) {
			return "pangram"
		}
	}
	return "not pangram"
}

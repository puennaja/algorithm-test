package palindome

import (
	"slices"
	"strings"
)

// check s is a  palindome string
// asbdbsa

func Palindome(s string) bool {
	str := strings.Split(s, "")
	slices.Reverse(str)
	newS := strings.Join(str, "")
	return s == newS
}

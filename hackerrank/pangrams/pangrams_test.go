package pangrams_test

import (
	"algorithm-test/hackerrank/pangrams"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyinteger(t *testing.T) {
	type caseTest struct {
		name     string
		args     string
		expected string
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     "We promptly judged antique ivory buckles for the prize",
			expected: "not pangrams",
		},
		{
			name: "case: 1",
			args: "qmExzBIJmdELxyOFWv LOCmefk TwPhargKSPEqSxzveiun",
			expected: "pangram",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := pangrams.Pangrams(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

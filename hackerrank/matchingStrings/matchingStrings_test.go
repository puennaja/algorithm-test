package matchingstrings_test

import (
	matchingstrings "algorithm-test/hackerrank/matchingStrings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchingStrings(t *testing.T) {
	type args struct {
		s []string
		q []string
	}
	type caseTest struct {
		name     string
		args     args
		expected []int32
	}

	cases := []caseTest{
		{
			name: "basic case: 1",
			args: args{
				s: []string{"aba", "baba", "aba", "xzxb"},
				q: []string{"aba", "xzxb", "ab", "k"},
			},
			expected: []int32{2, 1, 0, 0},
		},
		{
			name: "basic case: 2",
			args: args{
				s: []string{
					"abcde",
					"sdaklfj",
					"asdjf",
					"na",
					"basdn",
					"sdaklfj",
					"asdjf",
					"na",
					"asdjf",
					"na",
					"basdn",
					"sdaklfj",
					"asdjf",
				},
				q: []string{
					"abcde",
					"sdaklfj",
					"asdjf",
					"na",
					"basdn",
				},
			},
			expected: []int32{
				1,
				3,
				4,
				3,
				2,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := matchingstrings.MatchingStrings(tc.args.s, tc.args.q)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

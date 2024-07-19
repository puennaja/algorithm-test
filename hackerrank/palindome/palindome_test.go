package palindome_test

import (
	"algorithm-test/hackerrank/palindome"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindome(t *testing.T) {
	type caseTest struct {
		name     string
		args     string
		expected bool
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     "m",
			expected: true,
		},
		{
			name:     "basic case: 2",
			args:     "mxsd",
			expected: false,
		},
		{
			name:     "basic case: 3",
			args:     "racecar",
			expected: true,
		},
		{
			name:     "basic case: 4",
			args:     "wkdnofluwhbeouf",
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := palindome.Palindome(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

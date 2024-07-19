package timeconversion_test

import (
	"algorithm-test/hackerrank/timeconversion"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	type caseTest struct {
		name     string
		args     string
		expected string
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     "12:01:00AM",
			expected: "00:01:00",
		},
		{
			name:     "basic case: 2",
			args:     "12:01:00PM",
			expected: "12:01:00",
		},
		{
			name:     "basic case: 3",
			args:     "01:01:00AM",
			expected: "01:01:00",
		},
		{
			name:     "basic case: 4",
			args:     "01:01:00PM",
			expected: "13:01:00",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := timeconversion.TimeConversion(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

package countingsort_test

import (
	"algorithm-test/hackerrank/countingsort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	type caseTest struct {
		name     string
		args     []int32
		expected []int32
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     []int32{4, 3, 3, 2, 2, 1, 1},
			expected: []int32{1, 1, 2, 2, 3, 3, 4},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := countingsort.CountingSort(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

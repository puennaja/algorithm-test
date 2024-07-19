package bandprotocaltest_test

import (
	bandprotocaltest "algorithm-test/bandprotocal-test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupermanChickenRescue(t *testing.T) {
	type args struct {
		n int
		k int
		p []int
	}

	type caseTest struct {
		name     string
		args     args
		expected int
	}

	cases := []caseTest{
		{
			name: "basic case: 1",
			args: args{
				n: 5,
				k: 5,
				p: []int{2, 5, 10, 12, 15},
			},
			expected: 2,
		},
		{
			name: "basic case: 2",
			args: args{
				n: 6,
				k: 10,
				p: []int{1, 11, 30, 34, 35, 37},
			},
			expected: 4,
		},
		{
			name: "basic case: 3",
			args: args{
				n: 10,
				k: 5,
				p: []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11},
			},
			expected: 5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := bandprotocaltest.SupermanChickenRescue(tc.args.n, tc.args.k, tc.args.p)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

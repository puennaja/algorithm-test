package lonelyinteger_test

import (
	"algorithm-test/lonelyinteger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyinteger(t *testing.T) {
	type caseTest struct {
		name     string
		args     []int32
		expected int32
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     []int32{1, 1, 2, 2, 3, 3, 4},
			expected: int32(4),
		},
		{
			name:     "no lonely integer",
			args:     []int32{1, 1, 2, 2, 3, 3},
			expected: int32(-1),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lonelyinteger.Lonelyinteger(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

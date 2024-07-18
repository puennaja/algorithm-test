package bandprotocaltest_test

import (
	bandprotocaltest "algorithm-test/bandprotocal-test"
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
			args:     "SRSSRRR",
			expected: "Good Boy",
		},
		{
			name:     "basic case: 2",
			args:     "SSRSRRR",
			expected: "Good Boy",
		},
		{
			name:     "basic case: 3",
			args:     "RSSRR",
			expected: "Bad Boy",
		},
		{
			name:     "basic case: 4",
			args:     "SSSRRRS",
			expected: "Bad Boy",
		},
		{
			name:     "basic case: 5",
			args:     "SSSRRRRS",
			expected: "Bad Boy",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := bandprotocaltest.BossBabyRevenge(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

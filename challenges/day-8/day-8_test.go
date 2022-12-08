package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	tests := map[string]struct {
		input    string
		part2    bool
		expected int
	}{
		"part 1":   {input: "30373\n25512\n65332\n33549\n35390", part2: false, expected: 21},
		"part 1 b": {input: "30313\n25512\n65332\n33549\n35310", part2: false, expected: 22},
		"part 1 c": {input: "30313\n25512\n65332\n33559\n33549\n35310", part2: false, expected: 26},
		// "part 2":   {input: "30373\n25512\n65332\n33549\n35390", part2: true, expected: 8},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day8(tc.input, tc.part2))
		})
	}
}

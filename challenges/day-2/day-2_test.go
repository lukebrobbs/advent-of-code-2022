package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	tests := map[string]struct {
		input    string
		part2    bool
		expected int
	}{
		"part 1": {input: "A Y\nB X\nC Z", part2: false, expected: 15},
		"part 2": {input: "A Y\nB X\nC Z", part2: true, expected: 12},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day2(tc.input, tc.part2))
		})
	}
}

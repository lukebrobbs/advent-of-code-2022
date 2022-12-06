package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	tests := map[string]struct {
		input    string
		part2    bool
		expected int
	}{
		"part 1 a": {input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", part2: false, expected: 7},
		"part 1 b": {input: "bvwbjplbgvbhsrlpgdmjqwftvncz", part2: false, expected: 5},
		"part 1 c": {input: "nppdvjthqldpwncqszvftbrmjlhg", part2: false, expected: 6},
		"part 1 d": {input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", part2: false, expected: 10},
		"part 1 e": {input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", part2: false, expected: 11},

		"part 2 a": {input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", part2: true, expected: 19},
		"part 2 b": {input: "bvwbjplbgvbhsrlpgdmjqwftvncz", part2: true, expected: 23},
		"part 2 c": {input: "nppdvjthqldpwncqszvftbrmjlhg", part2: true, expected: 23},
		"part 2 d": {input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", part2: true, expected: 29},
		"part 2 e": {input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", part2: true, expected: 26},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day6(tc.input, tc.part2))
		})
	}
}

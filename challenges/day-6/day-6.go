package day6

import (
	"strings"
)

func Day6(input string, part2 bool) int {
	base := 3

	if part2 {
		base = 13
	}

	for i := range input {
		if i >= base {
			for ind, ch := range input[i-base : i+1] {
				if strings.Count(input[i-base:i+1], string(ch)) > 1 {
					break
				}
				if ind == base {
					return i + 1
				}
			}
		}
	}
	return 0
}

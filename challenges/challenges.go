package challenges

import (
	"github.com/lukebrobbs/advent-of-code-2022/challenges/day1"
)

type challenge func(input string, part2 bool) int

var DayChallenges = map[int]challenge{
	1: day1.Day1,
}

package challenges

import (
	day1 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-1"
	day2 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-2"
	day3 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-3"
	day4 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-4"
	day5 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-5"
	day6 "github.com/lukebrobbs/advent-of-code-2022/challenges/day-6"
)

type challenge func(input string, part2 bool) int

var DayChallenges = map[int]challenge{
	1: day1.Day1,
	2: day2.Day2,
	3: day3.Day3,
	4: day4.Day4,
	5: day5.Day5,
	6: day6.Day6,
}

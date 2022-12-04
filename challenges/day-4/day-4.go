package day4

import (
	"strconv"
	"strings"

	"github.com/lukebrobbs/advent-of-code-2022/utils"
)

func Day4(input string, part2 bool) (t int) {
	s := strings.Split(input, "\n")

	for _, pair := range s {
		p := strings.Split(pair, ",")
		first := createRange(p[0])
		second := createRange(p[1])
		shortest := len(first)
		if shortest > len(second) {
			shortest = len(second)
		}
		l := 0
		for _, f := range first {
			if utils.Contains(second, f) {
				l++
				if part2 {
					t++
					goto part2
				}
			}
		}

		if l == shortest {
			t++
		}

	part2:
		if l > 0 {
			continue
		}
	}
	return
}

func createRange(s string) (r []int) {
	p := strings.Split(s, "-")
	start, err := strconv.Atoi(p[0])
	utils.Check(err)
	end, err := strconv.Atoi(p[1])
	utils.Check(err)
	for i := start; i <= end; i++ {
		r = append(r, i)
	}
	return
}

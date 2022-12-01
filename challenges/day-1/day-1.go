package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/lukebrobbs/advent-of-code-2022/utils"
)

func Day1(input string, part2 bool) (o int) {
	values := strings.Split(input, "\n\n")
	totals := []int{}

	for _, v := range values {
		i := strings.Split(v, "\n")
		var t = 0
		for _, e := range i {
			st, err := strconv.Atoi(e)
			utils.Check(err)
			t += st
		}
		totals = append(totals, t)
	}
	// reverse a sorted slice
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	o = totals[0]
	if part2 {
		o += totals[1] + totals[2]
	}
	return
}

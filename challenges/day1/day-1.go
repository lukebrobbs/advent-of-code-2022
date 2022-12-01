package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Day1(input string, part2 bool) (b int) {
	values := strings.Split(input, "\n\n")
	totals := []int{}

	for _, v := range values {
		i := strings.Split(v, "\n")
		var t = 0
		for _, e := range i {
			st, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			t += st
		}
		totals = append(totals, t)
	}
	// reverse a sorted slice
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	b = totals[0]
	if part2 {
		b += totals[1] + totals[2]
	}
	return
}

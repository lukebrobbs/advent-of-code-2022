package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukebrobbs/advent-of-code-2022/utils"
)

type stack map[int][]string

func Day5(input string, part2 bool) int {
	i := strings.Split(input, "\n\n")
	sta := convertInputToStacks(i[0])
	instructions := createInstructions(i[1])

	for _, in := range instructions {
		for i := 1; i <= in[0]; i++ {
			from := sta[in[1]]
			to := sta[in[2]]
			if len(from) > 0 {
				sta[in[2]] = append([]string{from[0]}, to...)
				sta[in[1]] = from[1:]
			}
		}
	}
	str := ""

	for i := 1; i <= len(sta); i++ {
		if len(sta[i]) > 0 {
			str += sta[i][0]
		}
	}
	fmt.Print(str)
	return 5
}

func convertInputToStacks(input string) stack {
	s := stack{}

	for _, v := range strings.Split(input, "\n") {
		for i := 0; i < len(v); i += 4 {
			if string(v[i]) == "[" {
				in := ((i + 1) / 4) + 1
				s[in] = append(s[in], string(v[i+1]))
			}
		}
	}
	return s
}

func createInstructions(input string) [][]int {
	in := strings.Split(input, "\n")
	out := [][]int{}
	for _, i := range in {
		var move, from, to int
		cur := []int{}
		move = getIntFromString(i, "move ")
		from = getIntFromString(i, "from ")
		to = getIntFromString(i, "to ")
		cur = append(cur, move, from, to)
		out = append(out, cur)
	}
	return out
}

func getIntFromString(s, split string) int {
	t := strings.Split(s, split)[1]
	u := strings.Split(t, " ")
	ti, err := strconv.Atoi(string(u[0]))
	utils.Check(err)
	return ti
}

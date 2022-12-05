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
		if part2 {
			moveStacksPart2(in, sta)
		} else {
			moveStacks(in, sta)
		}
	}
	str := ""
	for i := 1; i <= len(sta); i++ {
		if len(sta[i]) > 0 {
			str += sta[i][0]
		}
	}
	fmt.Print(str)
	// add this to adhere to the interface needed to run the challenge from main.go
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

func moveStacks(in []int, sta stack) {
	for i := 1; i <= in[0]; i++ {
		from := sta[in[1]]
		to := sta[in[2]]
		if len(from) > 0 {
			sta[in[2]] = append([]string{from[0]}, to...)
			sta[in[1]] = from[1:]
		}
	}
}

func moveStacksPart2(in []int, sta stack) {
	var sl []string
	for i := 0; i < in[0]; i++ {
		sl = append(sl, sta[in[1]][0])
		if len(sta[in[1]]) > 1 {
			sta[in[1]] = sta[in[1]][1:]
		} else {
			sta[in[1]] = sta[in[1]][:0]
		}
	}
	sta[in[2]] = append(sl, sta[in[2]]...)
}

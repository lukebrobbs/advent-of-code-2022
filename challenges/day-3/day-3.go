package day3

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lukebrobbs/advent-of-code-2022/utils"
)

func Day3(input string, part2 bool) (t int) {
	bags := strings.Split(input, "\n")
	if part2 {
		t = handlePart2(bags)
		return
	}
	for _, b := range bags {
		matches := []rune{}
		len := utf8.RuneCountInString(b)
		b1 := b[0 : len/2]
		b2 := b[len/2 : len]
		for _, bi := range b1 {
			if strings.ContainsRune(b2, bi) {
				if !utils.RuneInSlice(bi, matches) {
					matches = append(matches, bi)
				}
			}
		}
		for _, m := range matches {
			t += getRuneVal(m)
		}
	}
	return
}

func getRuneVal(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	}
	return int(r) - 96
}

func handlePart2(bags []string) (t int) {
	g := []string{}
	for i, b := range bags {
		g = append(g, b)
		if (i+1)%3 == 0 {
			for _, s := range g[0] {
				if strings.ContainsRune(g[1], s) && strings.ContainsRune(g[2], s) {
					t += getRuneVal(s)
					break
				}
			}
			g = g[:0]
		}
	}
	return
}

package day3

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Day3(input string, part2 bool) (t int) {
	bags := strings.Split(input, "\n")
	t = 0
	for _, b := range bags {
		matches := []rune{}
		len := utf8.RuneCountInString(b)
		b1 := b[0 : len/2]
		b2 := b[len/2 : len]
		for _, bi := range b1 {
			if strings.ContainsRune(b2, bi) {
				if !runeInSlice(bi, matches) {
					matches = append(matches, bi)
				}
			}
		}
		for _, m := range matches {
			if unicode.IsUpper(m) {
				t += int(m) - 38
				continue
			}
			t += int(m) - 96
		}
	}
	return t
}

func runeInSlice(str rune, list []rune) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

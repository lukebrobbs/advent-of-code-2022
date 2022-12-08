package day8

import (
	"strings"
)

type Tree struct {
	height  rune
	visible bool
}

func Day8(input string, part2 bool) int {
	vis := 0
	tr := strings.Split(input, "\n")

	hor := checkHorizontal(tr)
	hor = checkVertical(hor)
	for _, h := range hor {
		for _, t := range h {
			if t.visible {
				vis++
			}
		}
	}
	return vis
}

func checkHorizontal(trees []string) [][]*Tree {
	vis := [][]*Tree{}
	for trI, tr := range trees {
		treeLs := []*Tree{}
		for i, t := range tr {
			if trI == 0 || trI == len(trees)-1 {
				treeLs = append(treeLs, &Tree{t, true})
				continue
			}
			visLeft, visRight := false, false
			if i == 0 || i == len(trees)-1 {
				treeLs = append(treeLs, &Tree{t, true})
				continue
			}
			for in, tree := range tr {
				if in == i {
					visLeft = true
					break
				}
				if tree >= t {
					break
				}

			}
			if visLeft {
				treeLs = append(treeLs, &Tree{t, true})
				continue
			}
			for tree := len(tr) - 1; tree >= i; tree-- {
				if tree == i {
					visRight = true
				}
				if rune(tr[tree]) >= t {
					break
				}
			}
			if visRight {
				treeLs = append(treeLs, &Tree{t, true})
				continue
			}
			treeLs = append(treeLs, &Tree{t, false})

		}
		vis = append(vis, treeLs)
	}
	return vis
}

func checkVertical(trees [][]*Tree) [][]*Tree {

	for trRow, row := range trees {
		for trI, tr := range row {

			if tr.visible {
				continue
			}
			for t := 0; t <= trRow; t++ {
				if t == trRow {
					tr.visible = true
					break
				}
				if trees[t][trI].height >= tr.height {
					break
				}
			}
			if tr.visible {
				continue
			}
			for t := trRow + 1; t <= len(trees)-1; t++ {
				if trees[t][trI].height >= tr.height {
					break
				}
				if t == len(trees)-1 {
					tr.visible = true
					break
				}
			}
		}
	}
	return trees
}

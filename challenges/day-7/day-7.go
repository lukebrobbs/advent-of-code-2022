package day7

import (
	"fmt"
	"path"
	"strings"
)

func Day7(input string, part2 bool) int {
	fs := map[string]int{}
	cd := ""
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		var size int
		var name string

		if strings.HasPrefix(s, "$ cd") {
			cd = path.Join(cd, strings.Fields(s)[2])
		} else if _, err := fmt.Sscanf(s, "%d %s", &size, &name); err == nil {
			for d := cd; d != "/"; d = path.Dir(d) {
				fs[d] += size
			}
			fs["/"] += size
		}

	}
	partA, partB := 0, fs["/"]
	for _, s := range fs {
		if s <= 100000 {
			partA += s
		}
		if s+70000000-fs["/"] >= 30000000 && s < partB {
			partB = s
		}
	}
	if part2 {
		return partB
	}
	return partA
}

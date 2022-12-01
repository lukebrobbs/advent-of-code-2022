package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/lukebrobbs/advent-of-code-2022/challenges"
)

var day = flag.Int("day", 1, "Advent of code day to run")
var part2 = flag.Bool("part-2", false, "Should we run part 2 of the code")

func main() {
	flag.Parse()
	challenges := challenges.DayChallenges
	bytes, err := ioutil.ReadFile(fmt.Sprintf("inputs/day-%d.txt", *day))
	if err != nil {
		panic(err)
	}
	fmt.Println(challenges[*day](string(bytes), *part2))
}

package day2

import (
	"strings"
)

type hand struct {
	value    int
	them     string
	me       string
	handType string
}

const ROCK = "rock"
const PAPER = "paper"
const SCISSORS = "scissors"

var outcomes = map[string]string{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

func Day2(input string, part2 bool) int {
	turns := strings.Split(input, "\n")

	points := []hand{
		{1, "A", "X", ROCK},
		{2, "B", "Y", PAPER},
		{3, "C", "Z", SCISSORS},
	}
	roundScore := 0
	for _, turn := range turns {
		t := strings.Split(turn, " ")
		var them, me hand
		for _, p := range points {
			if p.them == t[0] {
				them = p
			}
			if p.me == t[1] {
				me = p
			}
		}
		if part2 {
			p := gamePointsPart2(t[1], them)
			if p == 3 {
				roundScore += them.value
				roundScore += p
				continue
			}
			for _, po := range points {
				switch p {
				case 6:
					if outcomes[po.handType] == them.handType {
						roundScore += po.value
					}
				default:
					if po.handType == outcomes[them.handType] {
						roundScore += po.value
					}
				}
			}

			roundScore += p
			continue
		} else {
			p := gamePoints(them, me)
			roundScore += p
		}
	}

	return roundScore
}

func gamePoints(them, me hand) (o int) {
	o = me.value
	if outcomes[me.handType] == them.handType {
		o += 6
	}
	if them.handType == me.handType {
		o += 3
	}
	return
}

func gamePointsPart2(me string, them hand) int {
	points := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	return points[me]
}

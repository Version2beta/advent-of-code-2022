package main

import (
	"aoc/futils"
	"fmt"
	"strings"
)

func main() {
	lines := futils.Lines("../december-2-2022/input")
	fmt.Printf("Incorrect understanding: %d\n", RockPaperScissors(lines))
	fmt.Printf("Correct, strategic understanding: %d\n", StrategicRockPaperScissors(lines))
}

type round struct {
	us   string
	them string
}

type strategicRound struct {
	strategy string
	them     string
}

func RockPaperScissors(lines []string) int {
	score := 0
	for _, l := range lines {
		ps := strings.Split(l, " ")
		them := plays(ps[0])
		us := plays(ps[1])
		score = score + scores(us) + scores(outcome(round{us, them}))
	}
	return score
}

func StrategicRockPaperScissors(lines []string) int {
	score := 0
	for _, l := range lines {
		ps := strings.Split(l, " ")
		them := plays(ps[0])
		us := strategy(strategicRound{ps[1], them})
		score = score + scores(us) + scores(outcome(round{us, them}))
	}
	return score
}

func plays(k string) string {
	return map[string]string{
		"A": "rock",
		"X": "rock",
		"B": "paper",
		"Y": "paper",
		"C": "scissors",
		"Z": "scissors",
	}[k]
}

func scores(k string) int {
	return map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
		"lose":     0,
		"draw":     3,
		"win":      6,
	}[k]
}

func outcome(r round) string {
	return map[round]string{
		round{"rock", "scissors"}:     "win",
		round{"paper", "rock"}:        "win",
		round{"scissors", "paper"}:    "win",
		round{"rock", "paper"}:        "lose",
		round{"paper", "scissors"}:    "lose",
		round{"scissors", "rock"}:     "lose",
		round{"rock", "rock"}:         "draw",
		round{"paper", "paper"}:       "draw",
		round{"scissors", "scissors"}: "draw",
	}[r]
}

func strategy(r strategicRound) string {
	return map[strategicRound]string{
		strategicRound{"X", "rock"}:     "scissors", // lose
		strategicRound{"Y", "rock"}:     "rock",     // draw
		strategicRound{"Z", "rock"}:     "paper",    // win
		strategicRound{"X", "paper"}:    "rock",     // lose
		strategicRound{"Y", "paper"}:    "paper",    // draw
		strategicRound{"Z", "paper"}:    "scissors", // win
		strategicRound{"X", "scissors"}: "paper",    // lose
		strategicRound{"Y", "scissors"}: "scissors", // draw
		strategicRound{"Z", "scissors"}: "rock",     // win
	}[r]
}

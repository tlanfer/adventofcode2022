package main

import (
	"log"
	"strings"
)

func main() {
	rounds := strings.Split(input, "\n")

	totalScoreP1 := 0
	totalScoreP2 := 0

	for _, round := range rounds {
		hands := strings.Split(round, " ")
		opponent := toHand(hands[0])
		myself := toHand(hands[1])

		totalScoreP1 += playOneRound(opponent, myself) + int(myself)

		switch hands[1] {
		case "X": // need to lose
			myself = defeats(opponent)
		case "Y": // need to draw
			myself = opponent
		case "Z": // need to win
			myself = loses(opponent)
		}

		totalScoreP2 += playOneRound(opponent, myself) + int(myself)
	}

	log.Println("Total score P1:", totalScoreP1)
	log.Println("Total score P2:", totalScoreP2)
}

func playOneRound(opponent, myself Hand) int {
	if opponent == myself {
		return Draw
	}

	if defeats(myself) == opponent {
		return Win
	}

	return Lost
}

func toHand(s string) Hand {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}
	return 0
}

type Hand int

func defeats(h Hand) Hand {
	switch h {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	return 0
}

func loses(h Hand) Hand {
	switch h {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	return 0
}

const (
	Rock = iota + 1
	Paper
	Scissors
)

const (
	Lost = 0
	Draw = 3
	Win  = 6
)

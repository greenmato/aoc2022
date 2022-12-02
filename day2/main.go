package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
)

const (
	win  = "win"
	draw = "draw"
	lose = "lose"
)

// key beats value
var beats = map[string]string{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

// key loses to value
var loses = map[string]string{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

var choiceScores = map[string]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var outcomeScores = map[string]int{
	win:  6,
	draw: 3,
	lose: 0,
}

var part1Notation = map[string]string{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var part2Notation = map[string]string{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": lose,
	"Y": draw,
	"Z": win,
}

func hasWon(playerChoice string, opponentChoice string) bool {
	return beats[playerChoice] == opponentChoice
}

func getScore(playerChoice string, opponentChoice string) int {
	if playerChoice == opponentChoice {
		return outcomeScores[draw] + choiceScores[playerChoice]
	}
	if hasWon(playerChoice, opponentChoice) {
		return outcomeScores[win] + choiceScores[playerChoice]
	}

	return outcomeScores[lose] + choiceScores[playerChoice]
}

func getChoice(opponentChoice string, outcome string) string {
	if outcome == draw {
		return opponentChoice
	}
	if outcome == win {
		return loses[opponentChoice]
	}

	return beats[opponentChoice]
}

func parseChoices(line string, notation map[string]string) (string, string) {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		panic(errors.New("invalid input"))
	}

	return notation[split[0]], notation[split[1]]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	part1Score := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		opponentChoice, playerChoice := parseChoices(line, part1Notation)
		part1Score = part1Score + getScore(playerChoice, opponentChoice)
	}

	fmt.Printf("Part 1 score: %d\n", part1Score)

	part2Score := 0

	for _, line := range lines {
		opponentChoice, outcome := parseChoices(line, part2Notation)
		playerChoice := getChoice(opponentChoice, outcome)

		part2Score = part2Score + getScore(playerChoice, opponentChoice)
	}

	fmt.Printf("Part 2 score: %d\n", part2Score)
}

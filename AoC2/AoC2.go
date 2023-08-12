package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

type Round1 struct {
	OpponentShape Shape
	PlayerShape   Shape
}

type Round2 struct {
	OpponentShape  Shape
	DesiredOutcome int
}

func main() {
	lines := ReadInput("./input.txt")
	rounds1 := ParseInput1(lines)
	rounds2 := ParseInput2(lines)
	finalScore1 := Sol1(rounds1)
	finalScore2 := Sol2(rounds2)
	fmt.Println(finalScore1)
	fmt.Println(finalScore2)
}

func ReadInput(inputFilename string) []string {
	result := make([]string, 0)

	file, err := os.Open(inputFilename)
	if err != nil {
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func ParseInput1(input []string) []Round1 {
	result := make([]Round1, 0)
	for _, roundLine := range input {
		var (
			opponentShape Shape
			playerShape   Shape
		)
		shapes := strings.Split(roundLine, " ")
		switch shapes[0] {
		case "A":
			opponentShape = Rock
		case "B":
			opponentShape = Paper
		case "C":
			opponentShape = Scissors
		}

		switch shapes[1] {
		case "X":
			playerShape = Rock
		case "Y":
			playerShape = Paper
		case "Z":
			playerShape = Scissors
		}

		parsedRound := Round1{
			OpponentShape: opponentShape,
			PlayerShape:   playerShape,
		}

		result = append(result, parsedRound)
	}
	return result
}

func ParseInput2(input []string) []Round2 {
	result := make([]Round2, 0)
	for _, roundLine := range input {
		var (
			opponentShape  Shape
			desiredOutcome int
		)
		shapes := strings.Split(roundLine, " ")
		switch shapes[0] {
		case "A":
			opponentShape = Rock
		case "B":
			opponentShape = Paper
		case "C":
			opponentShape = Scissors
		}

		switch shapes[1] {
		case "X":
			desiredOutcome = 0
		case "Y":
			desiredOutcome = 3
		case "Z":
			desiredOutcome = 6
		}

		parsedRound := Round2{
			OpponentShape:  opponentShape,
			DesiredOutcome: desiredOutcome,
		}

		result = append(result, parsedRound)
	}
	return result
}

func Sol1(rounds []Round1) int {
	result := 0
	possibleOutcomes := [5]int{6, 0, 3, 6, 0}

	for _, round := range rounds {
		result += int(round.PlayerShape)
		roundDiff := round.PlayerShape - round.OpponentShape + 2
		result += possibleOutcomes[roundDiff]
	}
	return result
}

func Sol2(rounds []Round2) int {
	result := 0
	for _, round := range rounds {
		result += round.DesiredOutcome
		t := round.DesiredOutcome/3 - 1
		desiredShape := (int(round.OpponentShape) - 1 + t) % 3
		if desiredShape == -1 {
			desiredShape = 2
		}
		desiredShape = desiredShape + 1
		result += desiredShape
	}
	return result
}

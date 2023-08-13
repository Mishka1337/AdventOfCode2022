package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	From int
	To   int
}

type Pair struct {
	Assignments [2]Assignment
}

func main() {
	lines := ReadInput("./input.txt")
	pairs := ParseInput(lines)
	fmt.Println(Sol1(pairs))
	fmt.Println(Sol2(pairs))
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

func ParseInput(input []string) []Pair {
	res := make([]Pair, 0)
	for _, line := range input {
		parts := strings.Split(line, ",")
		pair := Pair{}
		for i, part := range parts {
			borders := strings.Split(part, "-")
			from, _ := strconv.ParseInt(borders[0], 10, 0)
			to, _ := strconv.ParseInt(borders[1], 10, 0)
			parsed := Assignment{
				From: int(from),
				To:   int(to),
			}
			pair.Assignments[i] = parsed
		}
		res = append(res, pair)
	}
	return res
}

func Sol1(input []Pair) int {
	res := 0

	for _, pair := range input {
		len1 := pair.Assignments[0].To - pair.Assignments[0].From + 1
		len2 := pair.Assignments[1].To - pair.Assignments[1].From + 1
		minRange := pair.Assignments[0]
		maxRange := pair.Assignments[1]
		if len1 > len2 {
			minRange = pair.Assignments[1]
			maxRange = pair.Assignments[0]
		}

		if minRange.From >= maxRange.From && minRange.To <= maxRange.To {
			res++
		}

	}

	return res
}

func Sol2(input []Pair) int {
	res := 0

	for _, pair := range input {
		lowerRange := pair.Assignments[0]
		higherRange := pair.Assignments[1]
		if lowerRange.From > higherRange.From {
			t := lowerRange
			lowerRange = higherRange
			higherRange = t
		}

		if lowerRange.To >= higherRange.From {
			res++
		}
	}

	return res
}

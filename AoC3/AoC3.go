package main

import (
	"bufio"
	"fmt"
	"os"
)

type ItemType int32

type Rucksack struct {
	FirstCompartment  []ItemType
	SecondCompartment []ItemType
}

func main() {
	lines := ReadInput("./input.txt")
	ruckSucks := ParseInput(lines)
	fmt.Println(Sol1(ruckSucks))
	fmt.Println(Sol2(ruckSucks))
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

func ParseInput(input []string) []Rucksack {
	result := make([]Rucksack, 0)

	for _, v := range input {
		s := []rune(v)
		lenS := len(s)
		firstHalf := make([]ItemType, 0)
		secondHalf := make([]ItemType, 0)

		for _, c := range s[:lenS/2] {
			var it ItemType
			if c >= 'a' && c <= 'z' {
				it = ItemType(c - 'a' + 1)
			}

			if c >= 'A' && c <= 'Z' {
				it = ItemType(c - 'A' + 27)
			}

			firstHalf = append(firstHalf, it)
		}

		for _, c := range s[lenS/2:] {
			var it ItemType
			if c >= 'a' && c <= 'z' {
				it = ItemType(c - 'a' + 1)
			}

			if c >= 'A' && c <= 'Z' {
				it = ItemType(c - 'A' + 27)
			}
			secondHalf = append(secondHalf, it)
		}

		ruckSuck := Rucksack{
			FirstCompartment:  firstHalf,
			SecondCompartment: secondHalf,
		}
		result = append(result, ruckSuck)
	}

	return result
}

func Sol1(rucksacks []Rucksack) int {
	result := 0
	for _, rs := range rucksacks {
		set := make(map[ItemType]struct{})
		for _, v := range rs.FirstCompartment {
			set[v] = struct{}{}
		}

		for _, v := range rs.SecondCompartment {
			if _, exists := set[v]; exists {
				result += int(v)
				break
			}
		}
	}
	return result
}

func Sol2(ruckSucks []Rucksack) int {
	result := 0
	for i := 2; i < len(ruckSucks); i += 3 {
		rs1 := append(ruckSucks[i-2].FirstCompartment, ruckSucks[i-2].SecondCompartment...)
		rs2 := append(ruckSucks[i-1].FirstCompartment, ruckSucks[i-1].SecondCompartment...)
		rs3 := append(ruckSucks[i].FirstCompartment, ruckSucks[i].SecondCompartment...)
		set1 := make(map[ItemType]struct{})
		set2 := make(map[ItemType]struct{})
		for _, v := range rs1 {
			set1[v] = struct{}{}
		}

		for _, v := range rs2 {
			if _, exists := set1[v]; exists {
				set2[v] = struct{}{}
			}
		}

		for _, v := range rs3 {
			if _, exists := set2[v]; exists {
				result += int(v)
				break
			}
		}
	}
	return result
}

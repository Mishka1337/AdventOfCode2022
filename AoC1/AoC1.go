package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	MealCal []int
}

func main() {
	lines := ReadInput("./input.txt")
	elves := ParseInput(lines)
	fmt.Println(Sol1(elves))
	fmt.Println(Sol2(elves))
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

func ParseInput(input []string) []Elf {
	result := make([]Elf, 0)

	curElf := Elf{
		MealCal: make([]int, 0),
	}
	for _, v := range input {
		if v == "" {
			result = append(result, curElf)
			curElf = Elf{
				MealCal: make([]int, 0),
			}
			continue
		}
		parsedCal, _ := strconv.ParseInt(v, 10, 0)
		curElf.MealCal = append(curElf.MealCal, int(parsedCal))
	}

	return result
}

func Sol1(elves []Elf) int {
	maxSum := 0
	for _, elf := range elves {
		sum := CalSum(elf)
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func Sol2(elves []Elf) int {
	top3Cal := [3]int{}
	top3Cal[0] = CalSum(elves[0])
	top3Cal[1] = CalSum(elves[1])
	top3Cal[2] = CalSum(elves[2])
	sort.Ints(top3Cal[:])

	for i := 3; i < len(elves); i++ {
		calSum := CalSum(elves[i])
		if calSum < top3Cal[0] {
			continue
		}
		InsertTop3(calSum, &top3Cal)
	}

	totalCal := 0
	for _, v := range top3Cal {
		totalCal += v
	}
	return totalCal
}

func CalSum(elf Elf) int {
	sumCal := 0
	for _, cal := range elf.MealCal {
		sumCal += cal
	}
	return sumCal
}

func InsertTop3(val int, arr *[3]int) {
	arr[0] = val
	for i := 1; i < 3; i++ {
		if arr[i-1] > arr[i] {
			temp := arr[i-1]
			arr[i-1] = arr[i]
			arr[i] = temp
		}
	}
}

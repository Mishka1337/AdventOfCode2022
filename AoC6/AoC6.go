package main

import (
	"fmt"
	"os"
)

func main() {
	input := ReadInput("./input.txt")
	fmt.Println(Sol1(input))
	fmt.Println(Sol2(input))
}

func ReadInput(fileName string) []rune {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return []rune{}
	}
	stringContent := string(fileContent)
	chars := []rune(stringContent)
	charsLen := len(chars)
	for i := 0; i < charsLen; i++ {
		chars[i] = chars[i] - 'a'
	}
	return chars
}

func Sol1(input []rune) int {
	charsLen := len(input)
	for i := 0; i < charsLen-3; {
		frame := input[i : i+4]
		res := FirstDup(frame)
		if res == -1 {
			return i + 4
		}
		i += res + 1
	}
	return charsLen
}

func Sol2(input []rune) int {
	charsLen := len(input)
	for i := 0; i < charsLen-13; {
		frame := input[i : i+14]
		res := FirstDup(frame)
		if res == -1 {
			return i + 14
		}
		i += res + 1
	}
	return charsLen
}

func FirstDup(chars []rune) int {
	bitVector := rune(0)
	dupBitVector := rune(0)
	for _, v := range chars {
		mask := rune(1 << v)
		if bitVector&mask == 0 {
			bitVector |= mask
		} else if dupBitVector&mask == 0 {
			dupBitVector |= mask
		}
	}

	for i, v := range chars {
		mask := rune(1 << v)
		if dupBitVector&mask != 0 {
			return i
		}
	}
	return -1
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AsmInstr interface {
	Exec(cpu *CPU)
}

type AddInstr struct {
	v int
}

type NoopInstr struct {
}

type CPU struct {
	X           int
	ClockCount  int
	Sol1Handler func(int, int)
	Sol2Handler func(int, int)
}

func (c *CPU) Incr() {
	c.ClockCount++
	if c.Sol2Handler != nil {
		c.Sol2Handler(c.X, c.ClockCount)
	}
	if c.ClockCount >= 20 && (c.ClockCount-20)%40 == 0 {
		if c.Sol1Handler != nil {
			c.Sol1Handler(c.X, c.ClockCount)
		}
	}
}

func (c *CPU) SubOnSol1Event(handler func(int, int)) {
	c.Sol1Handler = handler
}

func (c *CPU) SubOnSol2Event(handler func(int, int)) {
	c.Sol2Handler = handler
}

func (i AddInstr) Exec(cpu *CPU) {
	cpu.Incr()
	cpu.Incr()
	cpu.X += i.v
}

func (i NoopInstr) Exec(cpu *CPU) {
	cpu.Incr()
}

func main() {
	lines := ReadInput("./input.txt")
	asmCode := ParseInput(lines)
	fmt.Println(Sol1(asmCode))
	fmt.Println(Sol2(asmCode))
}

func ReadInput(filename string) []string {
	result := make([]string, 0)

	file, err := os.Open(filename)
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

func ParseInput(input []string) []AsmInstr {
	var res []AsmInstr

	for _, line := range input {
		if strings.HasPrefix(line, "noop") {
			res = append(res, NoopInstr{})
			continue
		}
		splited := strings.Split(line, " ")
		val, _ := strconv.Atoi(splited[1])
		res = append(res, AddInstr{v: val})
	}
	return res
}

func Sol1(asmCode []AsmInstr) int {
	res := 0
	cpu := CPU{X: 1}
	cpu.SubOnSol1Event(func(x int, clock int) {
		res += x * clock
	})

	for _, asm := range asmCode {
		asm.Exec(&cpu)
	}

	return res
}

func Sol2(asmCode []AsmInstr) string {
	resFormated := ""
	res := make([][]rune, 6)
	for i := range res {
		res[i] = make([]rune, 40)
	}

	cpu := CPU{X: 1}
	renderFinished := false
	cpu.SubOnSol2Event(func(x int, clock int) {
		index := clock - 1
		if index >= 240 {
			renderFinished = true
		}

		charToDraw := '.'
		curPixelCol := index % 40
		if curPixelCol >= x-1 && curPixelCol <= x+1 {
			charToDraw = '#'
		}

		row := index / 40
		col := index % 40
		res[row][col] = charToDraw
	})

	for _, asm := range asmCode {
		if renderFinished {
			break
		}
		asm.Exec(&cpu)
	}

	for _, runes := range res {
		line := string(runes)
		resFormated = fmt.Sprint(resFormated, line, "\n")
	}
	return resFormated
}

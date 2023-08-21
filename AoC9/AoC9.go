package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction byte

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type MoveCommand struct {
	Direction  Direction
	StepsCount int
}

type Point struct {
	X int
	Y int
}

type State struct {
	HeadKnotPos Point
	TailKnotPos Point
}

type StateLong struct {
	RopeKnotsPos [10]Point
}

func main() {
	lines := ReadInput("./input.txt")
	commands := ParseInput(lines)
	fmt.Println(Sol1(commands))
	fmt.Println(Sol2(commands))
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

func ParseInput(input []string) []MoveCommand {
	var res []MoveCommand
	for _, line := range input {
		splited := strings.Split(line, " ")
		steps, _ := strconv.Atoi(splited[1])
		command := MoveCommand{
			StepsCount: steps,
		}
		switch splited[0] {
		case "U":
			command.Direction = DirectionUp
		case "R":
			command.Direction = DirectionRight
		case "D":
			command.Direction = DirectionDown
		case "L":
			command.Direction = DirectionLeft
		}
		res = append(res, command)
	}
	return res
}

func Sol1(commands []MoveCommand) int {
	visitedPoints := make(map[Point]struct{})
	state := State{
		HeadKnotPos: Point{X: 0, Y: 0},
		TailKnotPos: Point{X: 0, Y: 0},
	}
	visitedPoints[state.TailKnotPos] = struct{}{}
	for _, c := range commands {
		for i := 0; i < c.StepsCount; i++ {
			state.HeadKnotPos.MakeMove(c.Direction)
			state.TailKnotPos.MoveTowards(state.HeadKnotPos)
			visitedPoints[state.TailKnotPos] = struct{}{}
		}
	}
	return len(visitedPoints)
}

func Sol2(commands []MoveCommand) int {
	visitedPoints := make(map[Point]struct{})
	state := StateLong{}
	visitedPoints[state.RopeKnotsPos[9]] = struct{}{}
	for _, c := range commands {
		for i := 0; i < c.StepsCount; i++ {
			state.RopeKnotsPos[0].MakeMove(c.Direction)
			for i := 1; i < 10; i++ {
				state.RopeKnotsPos[i].MoveTowards(state.RopeKnotsPos[i-1])
			}
			visitedPoints[state.RopeKnotsPos[9]] = struct{}{}
		}
	}
	return len(visitedPoints)
}

func (p *Point) MakeMove(dir Direction) {
	switch dir {
	case DirectionUp:
		p.Y += 1
	case DirectionRight:
		p.X += 1
	case DirectionDown:
		p.Y -= 1
	case DirectionLeft:
		p.X -= 1
	}
}

func (p *Point) MoveTowards(to Point) {
	if to.IsTouching(*p) {
		return
	}

	deltaX := to.X - p.X
	deltaY := to.Y - p.Y
	if deltaX < 0 {
		deltaX = -1
	} else if deltaX > 0 {
		deltaX = 1
	}
	if deltaY < 0 {
		deltaY = -1
	} else if deltaY > 0 {
		deltaY = 1
	}
	p.X += deltaX
	p.Y += deltaY
}

func (p Point) IsTouching(other Point) bool {
	deltaX := Abs(p.X - other.X)
	deltaY := Abs(p.Y - other.Y)
	if deltaX > 1 || deltaY > 1 {
		return false
	}
	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

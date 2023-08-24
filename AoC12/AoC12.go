package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Vector2 struct {
	X, Y int
}

type Point struct {
	Height        int
	CanReachedBy  []Vector2
	IsVisited     bool
	MinPathLength int
}

func main() {
	lines := ReadInput("./input.txt")
	points1, from1, to1 := ParseInput(lines)
	points2, _, to2 := ParseInput(lines)
	fmt.Println(Sol1(points1, from1, to1))
	fmt.Println(Sol2(points2, to2))
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

func ParseInput(input []string) ([][]Point, Vector2, Vector2) {
	var res [][]Point
	var from Vector2
	var to Vector2
	for y, line := range input {
		chars := []rune(line)
		var lineRes []Point
		for x, c := range chars {
			minPathLength := -1
			if c == 'E' {
				to = Vector2{
					X: x,
					Y: y,
				}
				c = 'z'
				minPathLength = 0
			}
			if c == 'S' {
				from = Vector2{
					X: x,
					Y: y,
				}
				c = 'a'
			}
			h := int(c - 'a')
			p := Point{
				Height:        h,
				MinPathLength: minPathLength,
			}
			lineRes = append(lineRes, p)
		}
		res = append(res, lineRes)
	}

	yMax := len(res)
	xMax := len(res[0])
	for y := range res {
		for x := range res[y] {
			for deltaY := -1; deltaY <= 1; deltaY++ {
				if deltaY == 0 {
					continue
				}
				yCur := y + deltaY
				if yCur < 0 || yCur >= yMax {
					continue
				}
				if res[y][x].Height-res[yCur][x].Height > 1 {
					continue
				}
				p := Vector2{
					X: x,
					Y: yCur,
				}
				res[y][x].CanReachedBy = append(res[y][x].CanReachedBy, p)
			}
			for deltaX := -1; deltaX <= 1; deltaX++ {
				if deltaX == 0 {
					continue
				}
				xCur := x + deltaX
				if xCur < 0 || xCur >= xMax {
					continue
				}
				if res[y][x].Height-res[y][xCur].Height > 1 {
					continue
				}
				p := Vector2{
					X: xCur,
					Y: y,
				}
				res[y][x].CanReachedBy = append(res[y][x].CanReachedBy, p)
			}
		}
	}

	return res, from, to
}

func Sol1(points [][]Point, from Vector2, to Vector2) int {
	for !points[from.Y][from.X].IsVisited {
		curIdx, _ := FindNotVisitedPointWithMinPath(points)
		curP := &points[curIdx.Y][curIdx.X]
		curLen := curP.MinPathLength
		for _, v := range curP.CanReachedBy {
			pathLen := curLen + 1
			nearPoint := &points[v.Y][v.X]
			if nearPoint.MinPathLength != -1 && nearPoint.MinPathLength < pathLen {
				continue
			}
			nearPoint.MinPathLength = pathLen
		}
		curP.IsVisited = true
	}
	return points[from.Y][from.X].MinPathLength
}

func Sol2(points [][]Point, to Vector2) int {
	var (
		ok     = true
		curIdx *Vector2
	)
	for ok {
		curIdx, ok = FindNotVisitedPointWithMinPath(points)
		if !ok {
			break
		}
		curP := &points[curIdx.Y][curIdx.X]
		curLen := curP.MinPathLength
		for _, v := range curP.CanReachedBy {
			pathLen := curLen + 1
			nearPoint := &points[v.Y][v.X]
			if nearPoint.MinPathLength != -1 && nearPoint.MinPathLength < pathLen {
				continue
			}
			nearPoint.MinPathLength = pathLen
		}
		curP.IsVisited = true
	}
	res := math.MaxInt
	for _, l := range points {
		for _, p := range l {
			if p.Height != 0 {
				continue
			}
			if p.MinPathLength != -1 && p.MinPathLength < res {
				res = p.MinPathLength
			}
		}
	}
	return res
}

func FindNotVisitedPointWithMinPath(points [][]Point) (*Vector2, bool) {
	min := math.MaxInt
	var res *Vector2
	for y, l := range points {
		for x, p := range l {
			if p.MinPathLength == -1 || p.IsVisited {
				continue
			}
			if p.MinPathLength < min {
				min = p.MinPathLength
				res = &Vector2{
					X: x,
					Y: y,
				}
			}
		}
	}
	ok := res != nil
	return res, ok
}

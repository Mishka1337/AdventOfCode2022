package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type CellType byte

const (
	Empty CellType = iota
	Rock
	Sand
)

type Vector2 struct {
	X, Y int
}

func (v Vector2) Add(other *Vector2) Vector2 {
	res := Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
	return res
}

type Vector2Iter struct {
	From, To Vector2
	Dest     *Vector2
}

func NewVector2Iterator(from, to Vector2) *Vector2Iter {
	return &Vector2Iter{
		From: from,
		To:   to,
	}
}

func (i *Vector2Iter) Next() bool {
	if i.Dest == nil {
		if i.From.X == i.To.X {
			if i.To.Y > i.From.Y {
				i.Dest = &Vector2{
					X: 0,
					Y: 1,
				}
			} else {
				i.Dest = &Vector2{
					X: 0,
					Y: -1,
				}
			}
		}
		if i.From.Y == i.To.Y {
			if i.To.X > i.From.X {
				i.Dest = &Vector2{
					X: 1,
					Y: 0,
				}
			} else {
				i.Dest = &Vector2{
					X: -1,
					Y: 0,
				}
			}
		}

		return true
	}

	res := i.From.X == i.To.X && i.From.Y == i.To.Y
	if res {
		return false
	}
	i.From = i.From.Add(i.Dest)
	return true
}

func (i *Vector2Iter) Get() Vector2 {
	return i.From
}

func main() {
	lines := ReadInput("./input.txt")
	caveMap1 := ParseInput(lines)
	caveMap2 := ParseInput(lines)
	fmt.Println(Sol1(caveMap1))
	fmt.Println(Sol2(caveMap2))
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

func ParseInput(input []string) map[int]map[int]CellType {
	res := make(map[int]map[int]CellType)
	var vecSeqs [][]Vector2
	for _, line := range input {
		vecsStr := strings.Split(line, " -> ")
		var vecsLine []Vector2
		for _, v := range vecsStr {
			vSplited := strings.Split(v, ",")
			x, _ := strconv.Atoi(vSplited[0])
			y, _ := strconv.Atoi(vSplited[1])
			vec := Vector2{
				X: x,
				Y: y,
			}
			vecsLine = append(vecsLine, vec)
		}
		vecSeqs = append(vecSeqs, vecsLine)
	}
	for _, vs := range vecSeqs {
		for i := 1; i < len(vs); i++ {
			vecIter := NewVector2Iterator(vs[i-1], vs[i])
			for vecIter.Next() {
				curr := vecIter.Get()
				xCurr := curr.X
				yCurr := curr.Y
				if _, ok := res[yCurr]; !ok {
					res[yCurr] = make(map[int]CellType)
				}
				res[yCurr][xCurr] = Rock
			}
		}
	}
	return res
}

func Sol1(caveMap map[int]map[int]CellType) int {
	maxDepth := -1
	for k := range caveMap {
		if k > maxDepth {
			maxDepth = k
		}
	}
	res := 0
	isFallenOut := false

	for !isFallenOut {
		newX := 500
		newY := 0
		isStill := false
		for !isStill {
			if newY > maxDepth {
				isFallenOut = true
				isStill = true
				continue
			}

			if caveMap[newY+1][newX] == Empty {
				newY++
				continue
			}

			if caveMap[newY+1][newX-1] == Empty {
				newY++
				newX--
				continue
			}

			if caveMap[newY+1][newX+1] == Empty {
				newY++
				newX++
				continue
			}
			if _, ok := caveMap[newY]; !ok {
				caveMap[newY] = make(map[int]CellType)
			}
			caveMap[newY][newX] = Sand
			isStill = true
			res++
		}
	}
	return res
}

func Sol2(caveMap map[int]map[int]CellType) int {
	maxDepth := -1
	for k := range caveMap {
		if k > maxDepth {
			maxDepth = k
		}
	}
	maxDepth += 2
	res := 0
	isSandSpawnerBlocked := false

	for !isSandSpawnerBlocked {
		newX := 500
		newY := 0
		isStill := false
		for !isStill {
			if newY+1 == maxDepth {
				isStill = true
				if _, ok := caveMap[newY]; !ok {
					caveMap[newY] = make(map[int]CellType)
				}
				caveMap[newY][newX] = Sand
				res++
				continue
			}

			if caveMap[newY+1][newX] == Empty {
				newY++
				continue
			}

			if caveMap[newY+1][newX-1] == Empty {
				newY++
				newX--
				continue
			}

			if caveMap[newY+1][newX+1] == Empty {
				newY++
				newX++
				continue
			}

			if newY == 0 {
				isSandSpawnerBlocked = true
				isStill = true
				res++
				continue
			}
			if _, ok := caveMap[newY]; !ok {
				caveMap[newY] = make(map[int]CellType)
			}
			caveMap[newY][newX] = Sand
			isStill = true
			res++
		}
	}
	return res
}

func VisualizeCaveMap(caveMap map[int]map[int]CellType) {
	minY := 0
	maxY := -1
	minX := math.MaxInt
	maxX := -1
	for k1, v1 := range caveMap {
		if k1 > maxY {
			maxY = k1
		}
		for k2 := range v1 {
			if k2 > maxX {
				maxX = k2
			}
			if k2 < minX {
				minX = k2
			}
		}
	}
	arrDepth := maxY - minY + 1
	arrWidth := maxX - minX + 1
	arr := make([][]rune, arrDepth)
	for i := range arr {
		arr[i] = make([]rune, arrWidth)
		for j := range arr[i] {
			if caveMap[i+minY][j+minX] == Rock {
				arr[i][j] = '#'
				continue
			}
			if caveMap[i+minY][j+minX] == Sand {
				arr[i][j] = 'o'
				continue
			}
			if caveMap[i+minY][j+minX] == Empty {
				arr[i][j] = '.'
				continue
			}
		}
		str := string(arr[i])
		fmt.Println(str)
	}
	time.Sleep(50 * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	Height      int
	IsVisible   bool
	ScenicScore int
}

func main() {
	lines := ReadInput("./input.txt")
	treeMap := ParseInput(lines)
	fmt.Println(Sol1(treeMap))
	fmt.Println(Sol2(treeMap))
}

func ReadInput(filename string) []string {
	var res []string

	file, err := os.Open(filename)
	if err != nil {
		return res
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func ParseInput(input []string) [][]Tree {
	var res [][]Tree

	for _, line := range input {
		chars := []rune(line)
		var lineRes []Tree
		for _, char := range chars {
			charString := string(char)
			height, _ := strconv.ParseInt(charString, 10, 0)
			tree := Tree{
				Height:      int(height),
				IsVisible:   false,
				ScenicScore: 0,
			}
			lineRes = append(lineRes, tree)
		}
		res = append(res, lineRes)
	}
	return res
}

func Sol1(treeMap [][]Tree) int {
	res := 0
	mapH := len(treeMap)
	mapW := len(treeMap[0])

	// from top
	for i := 0; i < mapH; i++ {
		preHeightMax := -1
		for j := 0; j < mapW; j++ {
			if treeMap[i][j].Height <= preHeightMax {
				continue
			}
			preHeightMax = treeMap[i][j].Height
			treeMap[i][j].IsVisible = true
			res++
		}
	}

	// from left
	for j := 0; j < mapW; j++ {
		preHeightMax := -1
		for i := 0; i < mapH; i++ {
			if treeMap[i][j].Height <= preHeightMax {
				continue
			}
			preHeightMax = treeMap[i][j].Height
			if treeMap[i][j].IsVisible {
				continue
			}
			treeMap[i][j].IsVisible = true
			res++
		}
	}

	// from right
	for j := 0; j < mapW; j++ {
		preHeightMax := -1
		for i := mapH - 1; i >= 0; i-- {
			if treeMap[i][j].Height <= preHeightMax {
				continue
			}
			preHeightMax = treeMap[i][j].Height
			if treeMap[i][j].IsVisible {
				continue
			}
			treeMap[i][j].IsVisible = true
			res++
		}
	}

	// from bottom
	for i := 0; i < mapH; i++ {
		preHeightMax := -1
		for j := mapW - 1; j >= 0; j-- {
			if treeMap[i][j].Height <= preHeightMax {
				continue
			}
			preHeightMax = treeMap[i][j].Height
			if treeMap[i][j].IsVisible {
				continue
			}
			treeMap[i][j].IsVisible = true
			res++
		}
	}

	return res
}

func Sol2(treeMap [][]Tree) int {
	res := 0
	mapH := len(treeMap)
	mapW := len(treeMap[0])

	for i := 0; i < mapH; i++ {
		for j := 0; j < mapW; j++ {
			CalcScenicScore(treeMap, i, j)
			tree := treeMap[i][j]
			if tree.ScenicScore > res {
				res = tree.ScenicScore
			}
		}
	}

	return res
}

func CalcScenicScore(treeMap [][]Tree, x int, y int) {
	curTree := &treeMap[x][y]
	mapH := len(treeMap)
	mapW := len(treeMap[0])

	//to the top
	for i := x - 1; i >= 0; i-- {
		t := treeMap[i][y]
		curTree.ScenicScore++
		if t.Height >= curTree.Height {
			break
		}
	}

	//to the left
	res := 0
	for i := y - 1; i >= 0; i-- {
		t := treeMap[x][i]
		res++
		if t.Height >= curTree.Height {
			break
		}
	}
	curTree.ScenicScore *= res

	//to the bottom
	res = 0
	for i := x + 1; i < mapH; i++ {
		t := treeMap[i][y]
		res++
		if t.Height >= curTree.Height {
			break
		}
	}
	curTree.ScenicScore *= res

	//to the right
	res = 0
	for i := y + 1; i < mapW; i++ {
		t := treeMap[x][i]
		res++
		if t.Height >= curTree.Height {
			break
		}
	}
	curTree.ScenicScore *= res
}

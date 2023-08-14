package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cargo rune

type CargoBox struct {
	Boxed Cargo
	Next  *CargoBox
}

type Command struct {
	From     int
	To       int
	Quantity int
}

func main() {
	lines := ReadInput("./test.txt")
	state1, coms1 := ParseInput(lines)
	state2, coms2 := ParseInput(lines)
	fmt.Println(Sol1(state1, coms1))
	fmt.Println(Sol2(state2, coms2))
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

func ParseInput(input []string) ([]*CargoBox, []Command) {
	resCom := make([]Command, 0)

	initStateLines := make([]string, 0)
	commandLines := make([]string, 0)
	for _, line := range input {
		if strings.HasPrefix(line, " 1") || line == "" {
			continue
		}

		if strings.HasPrefix(line, "m") {
			commandLines = append(commandLines, line)
		} else {
			initStateLines = append(initStateLines, line)
		}
	}

	resBoxLineLen := len([]rune(initStateLines[0]))
	resBoxLen := (resBoxLineLen + 1) / 4
	resBox := make([]*CargoBox, resBoxLen)

	for _, s := range initStateLines {
		chars := []rune(s)
		for i := 0; i < resBoxLen; i++ {
			linePos := i*4 + 1
			if chars[linePos] == ' ' {
				continue
			}
			resBox[i] = AppendEnd(resBox[i], Cargo(chars[linePos]))
		}
	}

	for _, s := range commandLines {
		p := strings.Split(s, " ")
		quant, _ := strconv.Atoi(p[1])
		from, _ := strconv.Atoi(p[3])
		to, _ := strconv.Atoi(p[5])

		command := Command{
			Quantity: quant,
			From:     from - 1,
			To:       to - 1,
		}
		resCom = append(resCom, command)
	}
	return resBox, resCom
}

func Sol1(state []*CargoBox, coms []Command) string {
	res := make([]rune, 0)

	for _, comm := range coms {
		for i := 0; i < comm.Quantity; i++ {
			orgn, dest := Move(state[comm.From], state[comm.To])
			state[comm.From] = orgn
			state[comm.To] = dest
		}
	}

	for _, st := range state {
		res = append(res, rune(st.Boxed))
	}

	return string(res)
}

func Sol2(state []*CargoBox, coms []Command) string {
	res := make([]rune, 0)

	for _, comm := range coms {
		orgn, dest := MoveBatch(state[comm.From], state[comm.To], comm.Quantity)
		state[comm.From] = orgn
		state[comm.To] = dest
	}

	for _, st := range state {
		res = append(res, rune(st.Boxed))
	}

	return string(res)
}

func AppendEnd(xs *CargoBox, x Cargo) *CargoBox {
	if xs == nil {
		xs = &CargoBox{
			Boxed: x,
			Next:  nil,
		}
		return xs
	}
	curCargoBox := xs
	for curCargoBox.Next != nil {
		curCargoBox = curCargoBox.Next
	}
	curCargoBox.Next = &CargoBox{
		Boxed: x,
		Next:  nil,
	}
	return xs
}

func Move(orgn *CargoBox, dest *CargoBox) (*CargoBox, *CargoBox) {
	t := orgn.Next
	orgn.Next = dest
	dest = orgn
	orgn = t
	return orgn, dest
}

func MoveBatch(orgn *CargoBox, dest *CargoBox, size int) (*CargoBox, *CargoBox) {
	orgnBeg := orgn
	orgnEnd := orgn
	for i := 1; i < size; i++ {
		if orgnEnd != nil {
			orgnEnd = orgnEnd.Next
		}
	}
	t := orgnEnd.Next
	orgnEnd.Next = dest
	dest = orgnBeg
	orgn = t
	return orgn, dest
}

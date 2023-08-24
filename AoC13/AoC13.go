package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PacketNode interface {
	fmt.Stringer
	CompareWith(PacketNode) int
}

type ListNode struct {
	Content []PacketNode
}

type IntNode struct {
	Content int
}

func (l *ListNode) String() string {
	res := "["
	for _, c := range l.Content {
		res = fmt.Sprintf("%s%s, ", res, c.String())
	}
	res = fmt.Sprintf("%s]", res)
	return res
}

func (i *IntNode) String() string {
	return fmt.Sprint(i.Content)
}

func (l *ListNode) CompareWith(n PacketNode) int {
	switch t := n.(type) {
	case *ListNode:
		curLen := len(l.Content)
		otherLen := len(t.Content)
		delta := curLen - otherLen

		minLen := curLen
		if curLen > otherLen {
			minLen = otherLen
		}

		for i := 0; i < minLen; i++ {
			curElem := l.Content[i]
			otherElem := t.Content[i]
			resElem := curElem.CompareWith(otherElem)
			if resElem != 0 {
				return resElem
			}
		}

		if delta == 0 {
			return 0
		}

		if delta < 0 {
			return -1
		}

		return 1

	case *IntNode:
		newNode := &ListNode{
			Content: []PacketNode{t},
		}
		return l.CompareWith(newNode)
	}
	return 0
}

func (i *IntNode) CompareWith(n PacketNode) int {
	switch t := n.(type) {
	case *IntNode:
		if i.Content > t.Content {
			return 1
		}

		if i.Content < t.Content {
			return -1
		}

		return 0
	case *ListNode:
		newNode := &ListNode{
			Content: []PacketNode{i},
		}
		return newNode.CompareWith(t)
	}
	return 0
}

func main() {
	lines := ReadInput("./input.txt")
	packetPairs := ParseInput1(lines)
	packets := ParseInput2(lines)
	fmt.Println(Sol1(packetPairs))
	fmt.Println(Sol2(packets))
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

func ParseInput1(input []string) [][2]PacketNode {
	var res [][2]PacketNode
	curPair := [2]PacketNode{}
	i := 0
	for _, line := range input {
		if line == "" {
			res = append(res, curPair)
			curPair = [2]PacketNode{}
			i = 0
			continue
		}
		curPair[i] = ParseNodeRec([]rune(line))
		i++
	}
	return res
}

func ParseInput2(input []string) []PacketNode {
	var res []PacketNode
	for _, line := range input {
		if line == "" {
			continue
		}
		parsed := ParseNodeRec([]rune(line))
		res = append(res, parsed)
	}
	return res
}

func ParseNodeRec(input []rune) PacketNode {
	var res PacketNode
	if input[0] == '[' {
		charsLen := len(input)
		resList := &ListNode{}

		for i := 1; i < charsLen-1; {
			nextPos := 1

			if input[i] >= '0' && input[i] <= '9' {
				nextPos = FindIntEnd(input[i:])
				newNode := ParseNodeRec(input[i : i+nextPos])
				resList.Content = append(resList.Content, newNode)
			}

			if input[i] == '[' {
				nextPos = FindListEnd(input[i:])
				newNode := ParseNodeRec(input[i : i+nextPos])
				resList.Content = append(resList.Content, newNode)
			}
			i += nextPos
		}
		res = resList
	} else {
		str := string(input)
		n, _ := strconv.Atoi(str)
		resInt := &IntNode{
			Content: n,
		}
		res = resInt
	}
	return res
}

func FindIntEnd(chars []rune) int {
	i := 0
	for chars[i] >= '0' && chars[i] <= '9' {
		i++
	}
	return i
}

func FindListEnd(chars []rune) int {
	i := 0
	depth := 0
	for depth > 0 || i == 0 {
		if chars[i] == '[' {
			depth++
		}
		if chars[i] == ']' {
			depth--
		}
		i++
	}
	return i
}

func Sol1(packetPairs [][2]PacketNode) int {
	var res int
	i := 1
	for _, pair := range packetPairs {
		cmpRes := pair[0].CompareWith(pair[1])
		if cmpRes != 1 {
			res += i
		}
		i++
	}
	return res
}

func Sol2(packets []PacketNode) int {
	divPackStr1 := "[[6]]"
	divPackStr2 := "[[2]]"
	divPack1 := ParseNodeRec([]rune(divPackStr1))
	divPack2 := ParseNodeRec([]rune(divPackStr2))
	packets = append(packets, divPack1, divPack2)
	packets = SortPackets(packets)
	idx1, idx2 := 0, 0
	for i, p := range packets {
		cmp1 := p.CompareWith(divPack1)
		cmp2 := p.CompareWith(divPack2)
		if cmp1 == 0 {
			idx1 = i + 1
		}
		if cmp2 == 0 {
			idx2 = i + 1
		}
	}
	return idx1 * idx2
}

// basically insertion sort, 'cause no need for something complex
func SortPackets(packets []PacketNode) []PacketNode {
	packLen := len(packets)
	for i := 1; i < packLen; i++ {
		key := packets[i]
		j := i - 1
		for j >= 0 && packets[j].CompareWith(key) == 1 {
			packets[j+1] = packets[j]
			j--
		}
		packets[j+1] = key
	}
	return packets
}

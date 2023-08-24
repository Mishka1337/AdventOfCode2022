package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MathOp int

const (
	MathOpAdd MathOp = iota
	MathOpMul
)

type Operation struct {
	OtherOperand *int
	MathOp       MathOp
}

func (o Operation) Perform1(val int) int {
	var x int
	if o.OtherOperand != nil {
		x = *o.OtherOperand
	} else {
		x = val
	}
	switch o.MathOp {
	case MathOpAdd:
		return x + val
	case MathOpMul:
		return x * val
	}
	return 0
}

func (o Operation) Perform2(val []int, divisors []int) []int {
	var res []int
	for i := range val {
		var x int
		if o.OtherOperand != nil {
			x = *o.OtherOperand
			x = x % divisors[i]
		} else {
			x = val[i]
		}

		switch o.MathOp {
		case MathOpAdd:
			v := (x + val[i]) % divisors[i]
			res = append(res, v)
		case MathOpMul:
			v := (x * val[i]) % divisors[i]
			res = append(res, v)
		}
	}
	return res
}

type Monkey struct {
	Items           []int
	Items2          [][]int //item value modulo for each monkey
	TestValue       int
	Operation       Operation
	FalseTestMonkey int
	TrueTestMonkey  int
	InspectionCount int
}

func main() {
	lines := ReadInput("./input.txt")
	ms1 := ParseInput(lines)
	ms2 := ParseInput(lines)
	fmt.Println(Sol1(ms1))
	fmt.Println(Sol2(ms2))
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

func ParseInput(input []string) []Monkey {
	var res []Monkey
	i := 0
	curMonkey := Monkey{}
	for _, line := range input {
		if line == "" {
			res = append(res, curMonkey)
			curMonkey = Monkey{}
			i = 0
			continue
		}
		line = strings.Trim(line, " ")
		lineSplited := strings.Split(line, " ")
		switch i {
		case 1:
			items := lineSplited[2:]
			for j := 0; j < len(items); j++ {
				s := DeleteNonNumberChars(items[j])
				num, _ := strconv.Atoi(s)
				curMonkey.Items = append(curMonkey.Items, num)
			}
		case 2:
			s1 := lineSplited[3]
			s2 := lineSplited[5]
			opS := lineSplited[4]
			switch opS {
			case "+":
				curMonkey.Operation.MathOp = MathOpAdd
			case "*":
				curMonkey.Operation.MathOp = MathOpMul
			}
			val1, err1 := strconv.Atoi(s1) // At least one will be with err
			val2, err2 := strconv.Atoi(s2)

			if err1 == nil {
				curMonkey.Operation.OtherOperand = &val1
			}
			if err2 == nil {
				curMonkey.Operation.OtherOperand = &val2
			}

		case 3:
			s := lineSplited[3]
			testVal, _ := strconv.Atoi(s)
			curMonkey.TestValue = testVal
		case 4:
			s := lineSplited[5]
			val, _ := strconv.Atoi(s)
			curMonkey.TrueTestMonkey = val
		case 5:
			s := lineSplited[5]
			val, _ := strconv.Atoi(s)
			curMonkey.FalseTestMonkey = val
		}
		i++
	}
	res = append(res, curMonkey)
	return res
}

func DeleteNonNumberChars(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if '0' <= b && b <= '9' {
			res.WriteByte(b)
		}
	}
	return res.String()
}

func Sol1(ms []Monkey) int {
	res := 0

	for i := 0; i < 20; i++ {
		for j := range ms {
			for _, item := range ms[j].Items {
				ms[j].InspectionCount++
				cur := ms[j].Operation.Perform1(item)
				toIdx := 0
				bored := math.Floor(float64(cur) / 3.0)
				cur = int(bored)
				testValue := ms[j].TestValue
				if cur%testValue == 0 {
					toIdx = ms[j].TrueTestMonkey
				} else {
					toIdx = ms[j].FalseTestMonkey
				}
				ms[toIdx].Items = append(ms[toIdx].Items, cur)
			}
			ms[j].Items = make([]int, 0)
		}
	}

	max1, max2 := FindTwoMaxInspectCount(ms)
	res = max1 * max2

	return res
}

// goofy af refactoring needed
func Sol2(ms []Monkey) int {
	res := 0
	var divisors []int

	for i := range ms {
		for j := range ms[i].Items {
			var res []int
			for k := range ms {
				v := ms[i].Items[j] % ms[k].TestValue
				res = append(res, v)
			}
			ms[i].Items2 = append(ms[i].Items2, res)
		}
		divisors = append(divisors, ms[i].TestValue)
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(ms); j++ {
			for k := range ms[j].Items2 {
				ms[j].InspectionCount++
				cur := ms[j].Operation.Perform2(ms[j].Items2[k], divisors)
				toIdx := 0
				if cur[j] == 0 {
					toIdx = ms[j].TrueTestMonkey
				} else {
					toIdx = ms[j].FalseTestMonkey
				}
				ms[toIdx].Items2 = append(ms[toIdx].Items2, cur)
			}
			ms[j].Items2 = make([][]int, 0)
		}
	}

	max1, max2 := FindTwoMaxInspectCount(ms)
	res = max1 * max2

	return res
}

func FindTwoMaxInspectCount(ms []Monkey) (int, int) {
	if len(ms) < 2 {
		return 0, 0
	}
	max1 := ms[0].InspectionCount
	max2 := ms[1].InspectionCount
	if max1 < max2 {
		t := max1
		max1 = max2
		max2 = t
	}
	for i := 2; i < len(ms); i++ {
		if ms[i].InspectionCount > max2 {
			max2 = ms[i].InspectionCount
		}
		if max1 < max2 {
			t := max1
			max1 = max2
			max2 = t
		}
	}
	return max1, max2
}

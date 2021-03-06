package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	instructions := strings.Split(string(input), "\n")
	max, supreme := largest(instructions)
	fmt.Printf("The largest value is '%v'\n", max)
	fmt.Printf("The highest value ever held was '%v'\n", supreme)
}

func largest(instructions []string) (max, supreme int) {
	registers := make(map[string]int)
	for _, cmd := range instructions {
		re := regexp.MustCompile(`\w+|[-0-9]+|[<=>!]+`)       // https://github.com/google/re2/wiki/Syntax
		t := re.FindAllStringSubmatch(cmd, -1)                // example cmd: av inc 640 if uea ==  0
		r, op, tr, cond := t[0][0], t[1][0], t[4][0], t[5][0] // vars:        r  op  inc -  tr cond test
		increment, _ := strconv.Atoi(string(t[2][0]))
		test, _ := strconv.Atoi(string(t[6][0]))
		if op == "dec" {
			increment *= -1
		}
		registers[r], _ = registers[r] // create register, if doesn't exist
		registers[tr], _ = registers[tr]
		value, _ := registers[tr]
		if (cond == ">" && value > test) || (cond == "<" && value < test) ||
			(cond == ">=" && value >= test) || (cond == "<=" && value <= test) ||
			(cond == "==" && value == test) || (cond == "!=" && value != test) {
			registers[r] += increment
		}
		if supreme < registers[r] {
			supreme = registers[r]
		}
	}
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return
}

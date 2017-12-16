package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)


type instruction struct {
	name string
	op string
	val int
}

func makeInstruction(s string) instruction {
	parts := strings.Split(s, " ")
	val, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		panic(err)
	}
	return instruction{name: parts[0], op: parts[1], val: int(val)}
}

func main() {
	file, err := os.Open("8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dict := make(map[string]int)
	maxVal := math.MinInt32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings := strings.Split(line, " if ")
		instr := makeInstruction(strings[0])
		condInstr := makeInstruction(strings[1])
		fail := false

		if _, ok := dict[condInstr.name]; !ok {
			dict[condInstr.name] = 0
		}

		if condInstr.op == "<" && !(dict[condInstr.name] < condInstr.val) {
			fail = true
		}
		if condInstr.op == "<=" && !(dict[condInstr.name] <= condInstr.val) {
			fail = true
		}
		if condInstr.op == ">" && !(dict[condInstr.name] > condInstr.val) {
			fail = true
		}
		if condInstr.op == ">=" && !(dict[condInstr.name] >= condInstr.val) {
			fail = true
		}
		if condInstr.op == "==" && !(dict[condInstr.name] == condInstr.val) {
			fail = true
		}
		if condInstr.op == "!=" && !(dict[condInstr.name] != condInstr.val) {
			fail = true
		}
		if fail {
			continue
		}

		if _, ok := dict[instr.name]; !ok {
			dict[instr.name] = 0
		}
		if instr.op == "inc" {
			dict[instr.name] += instr.val
		} else if instr.op == "dec" {
			dict[instr.name] -= instr.val
		}

		for _, val := range dict {
			maxVal = int(math.Max(float64(maxVal), float64(val)))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(maxVal)
}

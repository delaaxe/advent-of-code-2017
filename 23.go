package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	fmt.Println()
	file, err := os.Open("23.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var program [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		program = append(program, strs)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	execute2(program)
}

func execute2(program [][]string) {
	registers := map[string]int64{"a": 0}
	sum := 0
	index := 0
	for {
		if index == len(program) {
			break
		}
		//fmt.Println(index, registers)

		instruction := program[index]
		op := instruction[0]
		varX := instruction[1]
		varY := instruction[2]

		xValue, ok := registers[varX]
		if !ok {
			value, err := strconv.ParseInt(varX, 10, 32)
			if err != nil {
				//panic(err)
			}
			xValue = value
		}

		yValue, ok := registers[varY]
		if !ok {
			value, err := strconv.ParseInt(varY, 10, 32)
			if err != nil {
				panic(err)
			}
			yValue = value
		}

		if op == "set" {
			registers[varX] = yValue
		}
		if op == "add" {
			registers[varX] += yValue
		}
		if op == "sub" {
			registers[varX] -= yValue
		}
		if op == "mul" {
			sum += 1
			registers[varX] *= yValue
		}
		if op == "jnz" {
			if xValue != 0 {
				index += int(yValue)
				continue
			}
		}
		index++
	}
	fmt.Println(registers["a"], registers["b"], registers["c"], registers["d"], registers["e"], registers["f"], registers["g"], registers["h"])
	fmt.Println("sum", sum)
}

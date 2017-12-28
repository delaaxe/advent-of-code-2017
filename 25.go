package main

import (
	"fmt"
)

func main() {
	tape := make([]int, 100000)
	state := 'A'
	index := len(tape) / 2
	for i := 0; i < 12586542; i++ {
		val := tape[index]
		if state == 'A' {
			if val == 0 {
				tape[index] = 1
				index++
				state = 'B'
				continue
			}
			if val == 1 {
				tape[index] = 0
				index--
				state = 'B'
				continue
			}
			panic("")
		}
		if state == 'B' {
			if val == 0 {
				tape[index] = 0
				index++
				state = 'C'
				continue
			}
			if val == 1 {
				tape[index] = 1
				index--
				state = 'B'
				continue
			}
			panic("")
		}
		if state == 'C' {
			if val == 0 {
				tape[index] = 1
				index++
				state = 'D'
				continue
			}
			if val == 1 {
				tape[index] = 0
				index--
				state = 'A'
				continue
			}
			panic("")
		}
		if state == 'D' {
			if val == 0 {
				tape[index] = 1
				index--
				state = 'E'
				continue
			}
			if val == 1 {
				tape[index] = 1
				index--
				state = 'F'
				continue
			}
			panic("")
		}
		if state == 'E' {
			if val == 0 {
				tape[index] = 1
				index--
				state = 'A'
				continue
			}
			if val == 1 {
				tape[index] = 0
				index--
				state = 'D'
				continue
			}
			panic("")
		}
		if state == 'F' {
			if val == 0 {
				tape[index] = 1
				index++
				state = 'A'
				continue
			}
			if val == 1 {
				tape[index] = 1
				index--
				state = 'E'
				continue
			}
			panic("")
		}
		panic("")
	}
	sum := 0
	for _, val := range tape {
		sum += val
	}
	fmt.Println("max", sum)
}

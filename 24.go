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
	file, err := os.Open("24.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var values []int
		for _, value := range strings.Split(line, "/") {
			s, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				panic(err)
			}
			values = append(values, int(s))
		}
		ports = append(ports, values)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pin := 0
	maxlen, maxScore := -1, -1
	for _, startPort := range ports {
		if startPort[0] == 0 {
			pin = startPort[1]
		} else if startPort[1] == 0 {
			pin = startPort[0]
		} else {
			continue
		}
		results := findMatches(pin, [][]int{startPort}, ports)
		for _, result := range results {
			score := 0
			for _, port := range result {
				score += port[0] + port[1]
			}
			if len(result) >= maxlen {
				maxlen = len(result)
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	fmt.Println("max", maxScore)
}

func findMatches(pin int, state [][]int, ports [][]int) [][][]int {
	var results [][][]int
	for _, port := range ports {
		skip := false
		for _, statePort := range state {
			if statePort[0] == port[0] && statePort[1] == port[1] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		nextPin := -1
		if port[0] == pin {
			nextPin = port[1]
		} else if port[1] == pin {
			nextPin = port[0]
		}
		if nextPin != -1 {
			newState := make([][]int, len(state) + 1)
			copy(newState, state)
			newState[len(state)] = port

			results = append(results, newState)
			matches := findMatches(nextPin, newState, ports)
			for _, match := range matches {
				results = append(results, match)
			}
		}
	}
	return results
}
package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)

func main() {
	fmt.Println()
	file, err := os.Open("21.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := []string{".#.", "..#", "###"}


	var rules2x2 [][][]string
	var rules3x3 [][][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		pattern := strings.Split(parts[0], "/")
		replacement := strings.Split(parts[1], "/")
		rule := [][]string{pattern, replacement}
		if len(pattern) == 3 {
			rules3x3 = append(rules3x3, rule)
		} else {
			rules2x2 = append(rules2x2, rule)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(grid, "\n"))
	fmt.Println()

	for k := 0; k < 18; k++ {
		fmt.Println(k)
		if len(grid) % 2 == 0 {
			newGrid := make([]string, 3 * len(grid) / 2)
			for j := 0; j < len(grid); j += 2 {
				for i := 0; i < len(grid); i += 2 {
					didMatch := false
					subGrid := []string{
						grid[j][i:i+2],
						grid[j+1][i:i+2],
					}
					for _, rule := range rules2x2 {
						if matches2x2(subGrid, rule[0]) {
							didMatch = true
							repl := rule[1]
							newGrid[j/2*3] += repl[0]
							newGrid[j/2*3+1] += repl[1]
							newGrid[j/2*3+2] += repl[2]
							break
						}
					}
					if !didMatch {
						panic("didnt match")
					}
				}
			}
			grid = newGrid
			//fmt.Println(strings.Join(grid, "\n"))
			//fmt.Println()
		} else if len(grid) % 3 == 0 {
			newGrid := make([]string, 4 * len(grid) / 3)
			for j := 0; j < len(grid); j += 3 {
				for i := 0; i < len(grid); i += 3 {
					didMatch := false
					subGrid := []string{
						grid[j][i:i+3],
						grid[j+1][i:i+3],
						grid[j+2][i:i+3],
					}
					for _, rule := range rules3x3 {
						if matches3x3(subGrid, rule[0]) {
							didMatch = true
							repl := rule[1]
							newGrid[j/3*4] += repl[0]
							newGrid[j/3*4+1] += repl[1]
							newGrid[j/3*4+2] += repl[2]
							newGrid[j/3*4+3] += repl[3]
							break
						}
					}
					if !didMatch {
						panic("didnt match")
					}
				}
			}
			grid = newGrid
			//fmt.Println(strings.Join(grid, "\n"))
			//fmt.Println()
		} else {
			panic("not mod 2, 3")
		}
	}

	sum := 0
	for _, line := range grid {
		for _, char := range line {
			if char == '#' {
				sum += 1
			}
		}
	}
	fmt.Println("sum", sum)
}

/*

.#
..

..
.#

 */

func rotate2x2(g []string) []string {
	return []string{
		fmt.Sprintf("%c%c", g[1][0], g[0][0]),
		fmt.Sprintf("%c%c", g[1][1], g[0][1]),
	}
}

func matches2x2(g []string, pattern []string) bool {
	gs := strings.Join(g, "/")
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = []string{
		fmt.Sprintf("%c%c%c", pattern[0][1], pattern[0][0]),
		fmt.Sprintf("%c%c%c", pattern[1][1], pattern[1][0]),
	}
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate2x2(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	return false
}


func rotate3x3(g []string) []string {
	return []string{
		fmt.Sprintf("%c%c%c", g[2][0], g[1][0], g[0][0]),
		fmt.Sprintf("%c%c%c", g[2][1], g[1][1], g[0][1]),
		fmt.Sprintf("%c%c%c", g[2][2], g[1][2], g[0][2]),
	}
}

func matches3x3(g []string, pattern []string) bool {
	gs := strings.Join(g, "/")
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = []string{
		fmt.Sprintf("%c%c%c", pattern[0][2], pattern[0][1], pattern[0][0]),
		fmt.Sprintf("%c%c%c", pattern[1][2], pattern[1][1], pattern[1][0]),
		fmt.Sprintf("%c%c%c", pattern[2][2], pattern[2][1], pattern[2][0]),
	}
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	pattern = rotate3x3(pattern)
	if gs == strings.Join(pattern, "/") {
		return true
	}

	return false
}
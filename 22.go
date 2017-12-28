package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	fmt.Println()
	file, err := os.Open("22.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	padding := 1000
	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := make([]int, len(line)+2*padding)
		for i, value := range line {
			if value == '.' {
				values[i+padding] = 0
			} else if value == '#' {
				values[i+padding] = 1
			} else {
				panic("char")
			}
		}

		grid = append(grid, values)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < padding; i++ {
		line := make([]int, len(grid[0]))
		grid = append([][]int{line}, grid...)
		line = make([]int, len(grid[0]))
		grid = append(grid, line)
	}

	x, y := len(grid[0])/2, len(grid)/2
	dx, dy := 0, -1
	infections := 0
	for i := 0; i < 10000000; i++ {
		val := grid[y][x]
		if val == 1 { // turn right
			if dx == 0 {
				dx, dy = -dy, 0
			} else if dy == 0 {
				dx, dy = 0, dx
			} else {
				panic("")
			}
		} else if val == 0 { // turn left
			if dx == 0 {
				dx, dy = dy, 0
			} else if dy == 0 {
				dx, dy = 0, -dx
			} else {
				panic("")
			}
		} else if val == 3 {

		} else if val == 4 {
			if dx == 0 {
				dy = -dy
			} else if dy == 0 {
				dx = -dx
			} else {
				panic("")
			}
		} else {
			panic("")
		}

		if val == 0 {
			grid[y][x] = 3
		} else if val == 3 {
			grid[y][x] = 1
			infections++
		} else if val == 1 {
			grid[y][x] = 4
		} else if val == 4 {
			grid[y][x] = 0
		}

		x += dx
		y += dy
	}
	//for _, line := range grid {
		//fmt.Println(line)
	//}
	fmt.Println(infections)
}

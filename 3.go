package main

import (
	"fmt"
)

func main() {
	const input = 368078
	size := 3
	maxValue, grid := makeGrid(size)
	for maxValue < input {
		fmt.Println("grid size", size, "has max", maxValue)
		size += 2
		maxValue, grid = makeGrid(size)
	}
	printGrid(grid, size)
	fmt.Println("ok grid size", size, "has max", maxValue)
	var result int
	stop := false
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if grid[j][i] >= input {
				stop = true
				result = grid[j][i]
			}
		}
		if stop {
			break
		}
	}
	fmt.Println("result", result)
}

func makeGrid(size int) (int, [][]int) {
	x := 1
	step, lastStep := 0, 0
	direction := 'd'
	grid := make([][]int, size)
	for j := range grid {
		grid[j] = make([]int, size)
	}
	grid[(size-1)/2][(size-1)/2] = 1
	i := (size - 1) / 2
	j := (size-1)/2 - 1
	k := 1
	for {
		if direction == 'r' {
			i += 1
		} else if direction == 'u' {
			j -= 1
		} else if direction == 'l' {
			i -= 1
		} else if direction == 'd' {
			j += 1
		}
		if !(j >= 1 && j < size-1 && i >= 1 && i < size-1) {
			break
		}
		k -= 1;
		if k == 0 {
			if direction == 'r' {
				direction = 'u'
			} else if direction == 'u' {
				direction = 'l'
			} else if direction == 'l' {
				direction = 'd'
			} else if direction == 'd' {
				direction = 'r'
			}
			if step == lastStep {
				step += 1
			} else {
				lastStep = step
			}
			k = step
		}
		if i == (size-1)/2 && j == (size-1)/2 {
			x = 1
		} else {
			x = grid[j-1][i-1] + grid[j-1][i] + grid[j-1][i+1] + grid[j][i+1] + grid[j+1][i+1] + grid[j+1][i] + grid[j+1][i-1] + grid[j][i-1]
		}
		grid[j][i] = x
	}
	return x, grid
}

func printGrid(grid [][]int, size int) {
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			fmt.Printf("%8d", grid[j][i])
		}
		fmt.Println()
	}
}

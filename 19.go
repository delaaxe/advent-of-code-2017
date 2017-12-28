package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"unicode"
)

func main() {
	fmt.Println()
	file, err := os.Open("19.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	x, y := 0, 0
	for i, char := range grid[0] {
		if char == '|' {
			x = i
			fmt.Println(y, x)
		}
	}

	steps := 0
	width, height := len(grid[4]), len(grid)
	dx, dy := 0, 1
	var letters string
	for {
		steps++
		currentRune := grid[y][x]
		if currentRune == '+' {
			fmt.Println()
		}
		fmt.Println(y, x, string(currentRune))
		if unicode.IsLetter(currentRune) {
			letters += string(currentRune)
			fmt.Println(letters)
		}
		if y+dy < height && y+dy >= 0 && x+dx < width && x+dx >= 0 {
			nextRune := grid[y+dy][x+dx]
			if nextRune != ' ' {
				y += dy
				x += dx
				continue
			}
		}

		if dx == 0 {
			if x+1 < width && grid[y][x+1] != ' ' {
				dx, dy = 1, 0
			} else if x-1 >= 0 && grid[y][x-1] != ' ' {
				dx, dy = -1, 0
			} else {
				break
			}
		} else if dy == 0 {
			if y+1 < height && grid[y+1][x] != ' ' {
				dx, dy = 0, 1
			} else if y-1 >= 0 && grid[y-1][x] != ' ' {
				dx, dy = 0, -1
			} else {
				break
			}
		} else {
			panic("invalid")
		}
		y += dy
		x += dx
	}
	fmt.Println("steps", steps)
}

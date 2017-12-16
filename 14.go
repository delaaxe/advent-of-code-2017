package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "oundnydw"
	grid := [128][128]rune{}
	for i := 0; i < 128; i++ {
		hashInput := input + "-" + strconv.FormatInt(int64(i), 10)
		for j, hashValue := range hash(hashInput) {
			binaryStr := strconv.FormatInt(int64(hashValue), 2)
			for k, char := range fmt.Sprintf("%08s", binaryStr) {
				grid[i][j*8+k] = char
			}
		}
	}

	sum := 0
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			seen := map[string]bool{}
			walkGrid(x, y, grid, seen)
			if len(seen) > 0 {
				sum += 1
				for tag := range seen {
					x, y := tag2coords(tag)
					grid[y][x] = '0'
				}
			}
		}
	}
	fmt.Println(sum)
}

func walkGrid(x int, y int, grid [128][128]rune, seen map[string]bool) {
	if grid[y][x] == '0' {
		return
	}
	_, ok := seen[coords2tag(x, y)]
	if ok {
		return
	}
	seen[coords2tag(x, y)] = true
	if x > 0 && grid[y][x-1] == '1' {
		walkGrid(x-1, y, grid, seen)
	}
	if x < 127 && grid[y][x+1] == '1' {
		walkGrid(x+1, y, grid, seen)
	}
	if y > 0 && grid[y-1][x] == '1' {
		walkGrid(x, y-1, grid, seen)
	}
	if y < 127 && grid[y+1][x] == '1' {
		walkGrid(x, y+1, grid, seen)
	}
}

func coords2tag(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func tag2coords(tag string) (int, int) {
	parts := strings.Split(tag, ",")
	x, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		panic(err)
	}
	return int(x), int(y)
}

func hash(input string) []int {

	var sparseHash = make([]int, 256)
	var lengths []int

	for _, value := range input {
		lengths = append(lengths, int(value))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	for i := 0; i < 256; i++ {
		sparseHash[i] = i
	}

	index, skip := 0, 0
	for i := 0; i < 64; i++ {
		index, skip = hashRound(lengths, sparseHash, index, skip)
	}

	strHash := ""
	var denseHash = make([]int, 16)
	for i := 0; i < 16; i++ {
		res := sparseHash[i*16]
		for j := 1; j < 16; j++ {
			res = res ^ sparseHash[i*16+j]
		}
		denseHash[i] = res
		strHash += fmt.Sprintf("%x", res)
	}
	return denseHash
}

func hashRound(lengths []int, values []int, index int, skip int) (int, int) {
	for _, length := range lengths {
		if index+length < len(values) {
			selection := values[index: index+length]
			flip(selection)
			for i := index; i < index+length; i++ {
				values[i] = selection[i-index]
			}
		} else {
			selection := values[index:]
			selection = append(selection, values[:length-(len(values)-index)]...)
			flip(selection)
			j := 0
			for i := index; i < len(values); i, j = i+1, j+1 {
				values[i] = selection[j]
			}
			for i := 0; j < len(selection); i, j = i+1, j+1 {
				v := selection[j]
				values[i] = v
			}
		}
		index = (index + length + skip) % len(values)
		skip += 1
	}
	return index, skip
}

func flip(ss []int) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

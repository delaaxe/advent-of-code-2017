package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("11.txt")
	if err != nil {
		fmt.Println("error")
	}

	x, y, z := 0, 0, 0
	steps := strings.Split(string(body), ",")
	maxDistance := 0
	for _, step := range steps {
		switch step {
		case "n":
			x, y, z = x, y+1, z-1
		case "ne":
			x, y, z = x+1, y, z-1
		case "se\n":
			x, y, z = x+1, y-1, z
		case "se":
			x, y, z = x+1, y-1, z
		case "s":
			x, y, z = x, y-1, z+1
		case "sw":
			x, y, z = x-1, y, z+1
		case "nw":
			x, y, z = x-1, y+1, z
		default:
			panic(step)
		}
		distance := cube_distance([]int{0, 0, 0}, []int{x, y, z})
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println("dist:", maxDistance)
}

func cube_distance(a, b []int) int {
	return int(abs(a[0] - b[0]) + abs(a[1] - b[1]) + abs(a[2] - b[2])) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func contains(a []int, v int) bool {
	for _, b := range a {
		if v == b {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println()
	file, err := os.Open("20.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var particles [][]int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var values []int64
		for _, value := range strings.Split(line, ",") {
			s, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}
		particles = append(particles, values)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	removed := make([]int, len(particles))

	//distances := make([]float64, len(particles))
	for i := 0; i < 100000; i++ {
		for j, particle := range particles {
			if removed[j] == 1 {
				continue
			}
			p, v, a := particle[:3], particle[3:6], particle[6:]
			v[0] += a[0]
			v[1] += a[1]
			v[2] += a[2]
			p[0] += v[0]
			p[1] += v[1]
			p[2] += v[2]
		}

		for j := 0; j < len(particles); j++ {
			for k := j + 1; k < len(particles); k++ {
				p1 := particles[j][:3]
				p2 := particles[k][:3]
				if p1[0] == p2[0] && p1[1] == p2[1] && p1[2] == p2[2] {
					removed[j] = 1
					removed[k] = 1
				}
			}
		}

		sum := 0
		for _, value := range removed {
			if value == 0 {
				sum += 1
			}
		}
		fmt.Println(sum)

		/*		closest := -1
				minDist := math.MaxFloat64
				for i, particle := range particles {
					dist := math.Abs(float64(particle[0])) + math.Abs(float64(particle[1])) + math.Abs(float64(particle[2]))
					if dist < minDist {
						minDist = dist
						closest = i
					}
					distances[i] = dist
				}
				if i % 100000 == 0 {
					fmt.Println("closest", closest)
				}*/
	}
}

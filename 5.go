package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	var values []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic(err)
		}
		values = append(values, int(s))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := 0
	for {
		value := values[i]
		if i+value >= len(values) || i+value < 0 {
			sum += 1
			break
		}
		if value >= 3 {
			values[i] -= 1
		} else {
			values[i] += 1
		}
		i += value
		sum += 1
	}

	fmt.Println("sum:", sum)
}

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
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var values []int
		for _, value := range strings.Split(line, "\t") {
			s, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				panic(err)
			}
			values = append(values, int(s))
		}
		for i, value1 := range values {
			for _, value2 := range values[i+1:] {
				if value1 % value2 == 0 {
					sum += value1 / value2
					fmt.Println(value1 / value2, "=", value1, "/", value2)
				} else if value2 % value1 == 0 {
					sum += value2 / value1
					fmt.Println(value2 / value1, "=", value2, "/", value1)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum: %d", sum)
}

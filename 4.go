package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"sort"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		isValid := true
		values := strings.Split(line, " ")
		for i, value1 := range values {
			for _, value2 := range values[i+1:] {
				if areAnagrams(value1, value2) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		fmt.Println(isValid, line)
		if isValid {
			sum += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("sum:", sum)
}

func areAnagrams(a, b string) bool {
	values := strings.Split(a, "")
	sort.Strings(values)
	values2 := strings.Split(b, "")
	sort.Strings(values2)
	return strings.Join(values, "") == strings.Join(values2, "")
}
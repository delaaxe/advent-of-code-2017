package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	dict := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\w+) \(\d+\)(?: -> )?(.*)?`)
		match := re.FindStringSubmatch(line)
		parent := match[1]
		children := strings.Split(match[2], ", ")

		for _, child := range children {
			if child == "" {
				continue
			}
			dict[child] = parent
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for leaf := range dict {
		node := leaf
		for {
			if parent, ok := dict[node]; ok {
				//fmt.Println(node, "has parent", parent)
				node = parent
			} else {
				fmt.Println(node, "is root of leaf", leaf)
				break
			}
		}
	}
	fmt.Println("sum:", sum)
}

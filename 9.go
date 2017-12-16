package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("9.txt")
	if err != nil {
		fmt.Println("error")
	}

	stream := string(body)
	sum := 0

	for {
		i := strings.Index(stream, "!")
		if i == -1 {
			break
		}
		stream = stream[:i] + stream[i+2:]
	}

	for {
		i := strings.Index(stream, "<")
		j := strings.Index(stream, ">")
		if i == -1 || j == -i {
			break
		}
		if i >= j {
			panic("impossible state")
		}
		stream = stream[:i] + stream[j+1:]
	}

	stream = strings.Replace(stream, ",", "", -1)

	level := 0
	for _, r := range stream {
		c := string(r)
		if c == "{" {
			level += 1
			sum += level
		} else {
			level -= 1
		}
	}
	fmt.Println("sum:", sum)

	stream = string(body)
	sum = 0

	isGarbage := false
	garbageLength := 0
	for i := 0; i < len(stream); i++ {
		c := string(stream[i])
		if c == "<" && !isGarbage {
			isGarbage = true
			continue
		}
		if c == ">" && isGarbage {
			isGarbage = false
			sum += garbageLength
			garbageLength = 0
			continue
		}
		if c == "!" {
			i += 1
			continue
		}
		if isGarbage {
			garbageLength += 1
		}
	}
	fmt.Println("sum:", sum)
}

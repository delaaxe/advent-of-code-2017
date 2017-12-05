package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	body, err := ioutil.ReadFile("1.txt")
	if err != nil {
		fmt.Println("error")
	}
	var sum int
	for i := 0; i < len(body); i++ {
		a := body[i]
		b := body[(i + len(body)/2) % len(body)]
		if a == b {
			sum += int(a - '0')
			fmt.Printf("match at %d: %c %d %d\n", i, a, int(a - '0'), sum)
		}
	}
	fmt.Printf("sum: %d", sum)
}

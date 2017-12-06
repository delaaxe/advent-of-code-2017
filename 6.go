package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"math"
)

func main() {
	body, err := ioutil.ReadFile("6.txt")
	if err != nil {
		fmt.Println("error")
	}

	sum := 0
	var values []int
	var hashes []string

	for _, value := range strings.Split(string(body), "\t") {
		s, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			panic(err)
		}
		values = append(values, int(s))
	}

	for {
		var maxIndex int
		maxValue := math.MinInt32
		for i, value := range values {
			if value > maxValue {
				maxValue = value
				maxIndex = i
			}
		}
		values[maxIndex] = 0
		for maxValue > 0 {
			maxIndex = (maxIndex + 1) % len(values)
			values[maxIndex] += 1
			maxValue -= 1
		}
		newHash := makeHash(values)
		foundIndex := -1
		for i, hash := range hashes {
			if hash == newHash {
				foundIndex = i
				break
			}
		}
		sum += 1
		if foundIndex != -1 {
			fmt.Println("cycle:", len(hashes) - foundIndex)
			break
		} else {
			hashes = append(hashes, newHash)
		}
	}

	fmt.Println("sum:", sum)
}

func makeHash(ints []int) string {
	strs := make([]string, len(ints))
	for i, value := range ints {
		strs[i] = fmt.Sprintf("%2s", strconv.Itoa(value))
	}
	return strings.Join(strs, "|")
}

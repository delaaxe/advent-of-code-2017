package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	body, err := ioutil.ReadFile("10.txt")
	if err != nil {
		fmt.Println("error")
	}

	var sparseHash = make([]int, 256)
	var lengths []int

	fmt.Println(body)
	for _, value := range string(body) {
		lengths = append(lengths, int(value))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	for i := 0; i < 256; i++ {
		sparseHash[i] = i
	}

	index, skip := 0, 0
	for i := 0; i < 64; i++ {
		index, skip = round(lengths, sparseHash, index, skip)
	}

	var denseHash = make([]int, 16)
	for i := 0; i < 16; i++ {
		res := sparseHash[i*16]
		for j := 1; j < 16; j++ {
			res = res ^ sparseHash[i*16+j]
		}
		denseHash[i] = res
		fmt.Printf("%x", res)
	}
	fmt.Println()
}

func round(lengths []int, values []int, index int, skip int) (int, int) {
	for _, length := range lengths {
		if index+length < len(values) {
			selection := values[index: index+length]
			reverse(selection)
			for i := index; i < index+length; i++ {
				values[i] = selection[i-index]
			}
		} else {
			selection := values[index:]
			selection = append(selection, values[:length-(len(values)-index)]...)
			reverse(selection)
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

func reverse(ss []int) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

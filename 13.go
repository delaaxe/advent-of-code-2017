package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

type layer struct {
	depth int
	scannerIndex int
	scannerDirection int
}

func main() {
	fmt.Println()
	file, err := os.Open("13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lastLayerIndex := math.MinInt32
	dict := map[int]layer{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ": ")
		layerIndex, err := strconv.ParseInt(strs[0], 10, 32)
		if err != nil {
			panic(err)
		}
		if int(layerIndex) > lastLayerIndex {
			lastLayerIndex = int(layerIndex)
		}
		depth, err := strconv.ParseInt(strs[1], 10, 32)
		if err != nil {
			panic(err)
		}
		dict[int(layerIndex)] = layer{
			depth:int(depth),
			scannerIndex:0,
			scannerDirection:-1,
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for passed, i := false, 1; !passed; i++ {
		if i % 100000 == 0 {
			fmt.Println(i)
		}
		step(dict)
		d := map[int]layer{}
		for key, val := range dict {
			d[key] = val
		}
		passed = cross(lastLayerIndex, d)
		if passed {
			fmt.Println("passed after delay", i)
		}
	}
}

func cross(lastLayerIndex int, dict map[int]layer) bool {
	for i := 0; i <= lastLayerIndex; i++ {
		layer, ok := dict[i]
		//fmt.Println(i, dict)
		if ok {
			if layer.scannerIndex == 0 {
				//fmt.Println("caught at", i)
				return false
			}
		}
		step(dict)
	}
	return true
}

func step(dict map[int]layer) {
	for key, layer := range dict {
		if layer.scannerIndex == 0 {
			layer.scannerIndex = 1
			layer.scannerDirection = 1
		} else if layer.scannerIndex == layer.depth-1 {
			layer.scannerIndex = layer.depth - 2
			layer.scannerDirection = -1
		} else {
			layer.scannerIndex += layer.scannerDirection
		}
		dict[key] = layer
	}
}

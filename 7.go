package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	childToParent := map[string]string{}
	parentToChildren := map[string][]string{}
	individualWeights := map[string]int{}
	totalWeights := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\w+) \((\d+)\)(?: -> )?(.*)?`)
		match := re.FindStringSubmatch(line)
		parent := match[1]
		weight, err := strconv.ParseInt(match[2], 10, 32)
		if err != nil {
			panic(err)
		}
		individualWeights[parent] = int(weight)
		children := strings.Split(match[3], ", ")
		var validChildren []string

		for _, child := range children {
			if child == "" {
				continue
			}
			childToParent[child] = parent
			validChildren = append(validChildren, child)
		}
		parentToChildren[parent] = validChildren
		if len(children) == 0 {
			totalWeights[parent] = int(weight)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for leaf := range childToParent {
		node := leaf
		for {
			if parent, ok := childToParent[node]; ok {
				node = parent
			} else {
				//fmt.Println(node, "is root of leaf", leaf)
				break
			}
		}
	}
	fmt.Println("sum:", sum)

	/*
	for _, node := range childToParent {
		fmt.Print(node, ": ")
		for child, parent := range childToParent {
			if parent == node {
				fmt.Print(individualWeights[child], " ")
			}
		}
		fmt.Println()
	}
	*/
	_, totalWeights = getTotalWeight("cqmvs", totalWeights, individualWeights, parentToChildren)
	fmt.Println("bntzksk", totalWeights["bntzksk"])
	fmt.Println("mvpqv", totalWeights["mvpqv"])
	fmt.Println("znztzxd", totalWeights["znztzxd"])
	fmt.Println()
	fmt.Println("vmttcwe", individualWeights["vmttcwe"], totalWeights["vmttcwe"])
	fmt.Println("ukwlfcf", individualWeights["ukwlfcf"], totalWeights["ukwlfcf"])
	fmt.Println("zzpevgd", individualWeights["zzpevgd"], totalWeights["zzpevgd"])
}

func getTotalWeight(node string, totalWeights map[string]int, individualWeights map[string]int,
	parentToChildren map[string][]string) (int, map[string]int) {
	if weight, ok := totalWeights[node]; ok {
		return weight, totalWeights
	}
	individualWeight, ok := individualWeights[node]
	if !ok {
		panic("no weight for " + node)
	}
	total := individualWeight
	lastWeight := -1
	children := parentToChildren[node]
	for _, child := range children {
		totalWeight, totalWeightsCopy := getTotalWeight(child, totalWeights, individualWeights, parentToChildren)
		if lastWeight == -1 {
			lastWeight = totalWeight
		} else if lastWeight != totalWeight {
			fmt.Println("different child weight for parent", node, individualWeight, "with children", children)
		}
		total += totalWeight
		totalWeights = totalWeightsCopy
	}
	totalWeights[node] = total
	return total, totalWeights
}

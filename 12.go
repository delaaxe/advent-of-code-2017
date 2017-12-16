package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

func main() {
	file, err := os.Open("12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dict := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " <-> ")
		source, err := strconv.ParseInt(strs[0], 10, 32)
		if err != nil {
			panic(err)
		}
		var targets []int
		for _, targetString := range strings.Split(strs[1], ", ") {
			target, err := strconv.ParseInt(targetString, 10, 32)
			if err != nil {
				panic(err)
			}
			targets = append(targets, int(target))
		}
		dict[int(source)] = targets
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	groups := make(map[string]bool)
	for source := range dict {
		seen := make(map[int]bool)
		walkGraph(source, dict, seen)

		keys := make([]int, len(seen))
		i := 0
		for key := range seen {
			keys[i] = key
			i++
		}
		sort.Ints(keys)
		group := ""
		for _, key := range keys {
			group += strconv.FormatInt(int64(key), 10) + "/"
		}
		fmt.Println(source, group)
		groups[group] = true
	}
	fmt.Println(len(groups))
}

func walkGraph(source int, dict map[int][]int, seen map[int]bool) {
	for _, target := range dict[source] {
		_, alreadySeen := seen[target]
		if alreadySeen {
			continue
		}
		seen[target] = true
		walkGraph(target, dict, seen)
	}
}
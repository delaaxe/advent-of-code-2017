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
	fmt.Println()
	file, err := os.Open("18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var program [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		program = append(program, strs)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	channel0 := make(chan int64, 9999999)
	channel1 := make(chan int64, 9999999)
	done := make(chan bool, 2)
	go execute(program, channel0, channel1, done, 0)
	go execute(program, channel1, channel0, done, 1)
	<-done
	<-done
}

func execute(program [][]string, sndChannel chan<- int64, rcvChannel <-chan int64, done chan<- bool, programId int) {
	registers := map[string]int64{"p": int64(programId)}
	sends := 0
	index := int64(0)
	for {
		instruction := program[index]
		op := instruction[0]
		varX := instruction[1]
		if op == "snd" {
			sends++
			if programId == 1 {
				fmt.Println("***", programId, "sending", sends, "value", registers[varX])
			}
			sndChannel <- registers[varX]
			index++
			continue
		} else if op == "rcv" {
			registers[varX] = <-rcvChannel
			index++
			continue
		}
		varY := instruction[2]
		yValue, ok := registers[varY]
		if !ok {
			value, err := strconv.ParseInt(varY, 10, 32)
			if err != nil {
				panic(err)
			}
			yValue = value
		}
		if op == "set" {
			registers[varX] = yValue
		}
		if op == "add" {
			registers[varX] += yValue
		}
		if op == "mul" {
			registers[varX] *= yValue
		}
		if op == "mod" {
			registers[varX] %= yValue
		}
		if op == "jgz" {
			xValue, ok := registers[varX]
			if !ok {
				value, err := strconv.ParseInt(varX, 10, 32)
				if err != nil {
					panic(err)
				}
				xValue = value
			}
			if xValue > 0 {
				index += yValue
				continue
			}
		}
		index++
	}
	fmt.Println(programId, "sent", sends, "times")
	done <- true
}

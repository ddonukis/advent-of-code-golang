package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const bankCount = 16

func main() {
	if argc := len(os.Args); argc != 2 {
		log.Fatalf("Exepcted exactly 1 arg (path to input file), received: %d\n", argc)
	}
	path := os.Args[1]
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	var numbers [bankCount]int

	for i, s := range fields {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("bad number '%s': %v\n", s, err)
		}
		numbers[i] = val
	}

	steps, loopLen := findCycle(numbers)
	fmt.Printf("Part 1: %d\n", steps)
	fmt.Printf("Part 2: %d\n", loopLen)
}

func findCycle(banks [bankCount]int) (int, int) {
	cycles := 0
	loopLen := 0
	states := make(map[[bankCount]int]int)
	states[banks] = 0

	for {
		cycles++
		nextBankId := mostBlocksBankIndex(&banks)
		blocks := banks[nextBankId]
		banks[nextBankId] = 0
		for i := 1; i <= blocks; i++ {
			idx := (nextBankId + i) % bankCount
			banks[idx] += 1
		}
		if c, known := states[banks]; known {
			loopLen = cycles - c
			break
		} else {
			states[banks] = cycles
		}
	}
	return cycles, loopLen
}

func mostBlocksBankIndex(banks *[bankCount]int) int {
	maxVal := math.MinInt
	maxValId := -1

	for i, blockCount := range banks {
		if blockCount > maxVal {
			maxVal = blockCount
			maxValId = i
		}
	}
	return maxValId
}

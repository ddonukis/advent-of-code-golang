package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		nums := processLine(line)

		totalP1 := 0
		totalP2 := 0
		offsetP1 := 1
		offsetP2 := len(nums) / 2
		for i, n := range nums {
			if n == nums[(i+offsetP1)%len(nums)] {
				totalP1 += n
			}
			if n == nums[(i+offsetP2)%len(nums)] {
				totalP2 += n
			}
		}
		fmt.Printf("Part 1: %d\n", totalP1)
		fmt.Printf("Part 2: %d\n", totalP2)
	}
}

func processLine(line string) []int {
	nums := make([]int, len(line))

	for i, ch := range line {
		val, err := strconv.Atoi(string([]rune{ch}))
		if err != nil {
			log.Fatalf("bad digit '%c'\n", ch)
		}
		nums[i] = val
	}
	return nums
}

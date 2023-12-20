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
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	jumpOffsets := make([]int, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()

		offset, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("bad line '%s': %v\n", line, err)
		}
		jumpOffsets = append(jumpOffsets, offset)
	}

	count := jump(jumpOffsets)

	fmt.Printf("Part 2: %d\n", count)
}

func jump(offsets []int) int {
	stepCount := 0
	nextInstructionIdx := 0

	for nextInstructionIdx < len(offsets) && nextInstructionIdx >= 0 {
		n := offsets[nextInstructionIdx]
		if n >= 3 {
			offsets[nextInstructionIdx] -= 1
		} else {
			offsets[nextInstructionIdx] += 1
		}

		nextInstructionIdx += n
		stepCount++
	}

	return stepCount
}

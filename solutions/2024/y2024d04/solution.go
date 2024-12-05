package y2024d04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Printf("AoC 2024 - day 3\n\n")

	charMatrix, err := readCharMatrix(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	t0 := time.Now()
	result1 := Part1(charMatrix)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(charMatrix)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func readCharMatrix(inputPath string) (charMatrix CharMatrix, err error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	charMatrix = make(CharMatrix, 0)
	for scanner.Scan() {
		line := scanner.Text()
		charMatrix = append(charMatrix, []byte(line))
	}
	return
}

type CharMatrix [][]byte

func (m CharMatrix) String() string {
	capacity := 0
	for _, ln := range m {
		capacity += len(ln)
	}

	elements := make([]string, 0, capacity)
	for _, ln := range m {
		for _, ch := range ln {
			elements = append(elements, fmt.Sprintf("%c", ch))
		}
		elements = append(elements, fmt.Sprintln())
	}
	return strings.Join(elements, "")
}

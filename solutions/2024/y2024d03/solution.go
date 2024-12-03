package y2024d03

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Printf("AoC 2024 - day 3\n\n")

	t0 := time.Now()
	result1 := Part1(inputPath)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(inputPath)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func readInput(inputPath string) (content string, err error) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return "", err
	}
	content = string(input)
	return
}

func productOfMul(input string) int {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	found := pattern.FindAllStringSubmatch(input, -1)

	fmt.Printf("%d matches found\n", len(found))

	sum := 0
	for _, match := range found {
		product := 1
		for _, item := range match[1:] {
			operand, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalln(err)
			}
			product *= operand
		}
		sum += product
	}
	return sum
}

func Part1(inputPath string) int {
	inputStr, err := readInput(inputPath)
	if err != nil {
		log.Fatalln(err)
	}
	return productOfMul(inputStr)
}

func Part2(inputPath string) int {
	inputStr, err := readInput(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	sum := 0

	cursor := 0
	enabled := true
	var nextSearchKeyword string
	for {
		if enabled {
			nextSearchKeyword = "don't()"
		} else {
			nextSearchKeyword = "do()"
		}

		nextCursor := strings.Index(inputStr[cursor:], nextSearchKeyword)

		if nextCursor == -1 && !enabled {
			break
		} else if nextCursor == -1 {
			sum += productOfMul(inputStr[cursor:])
			break
		}

		nextCursor += len(inputStr[:cursor])

		if enabled {
			sum += productOfMul(inputStr[cursor:nextCursor])
			enabled = false
		} else {
			enabled = true
		}

		cursor = nextCursor + len(nextSearchKeyword)
	}

	return sum

}

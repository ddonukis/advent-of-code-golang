package y2024d02

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(inputPath string) {
	fmt.Println("AOC 2024 - day 02")

	part1Result := Part1(inputPath)
	fmt.Printf("Part 1: %d\n", part1Result)

	part2Result := Part2(inputPath)
	fmt.Printf("Part 2: %d\n", part2Result)
}

func Part1(dataFilePath string) int {

	file, err := os.Open(dataFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var graduallyChaningCount int
	for scanner.Scan() {
		line := scanner.Text()

		numbers, err := parseLine(line)
		if err != nil {
			log.Fatalln(err)
		}

		if isGraduallyChanging(numbers) {
			graduallyChaningCount += 1
		}
	}
	return graduallyChaningCount
}

func isGraduallyChanging(numbers []int) bool {
	var isIncreasing bool
	for i, n := range numbers[1:] {
		diff := n - numbers[i]
		if diff == 0 {
			return false
		}
		if i == 0 && diff > 0 {
			isIncreasing = true
		}
		if isIncreasing && (diff < 1 || diff > 3) {
			return false
		}
		if !isIncreasing && (diff < -3 || diff > -1) {
			return false
		}
	}
	return true
}

func Part2(dataFilePath string) int {
	return 0
}

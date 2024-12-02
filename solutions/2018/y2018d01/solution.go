package y2018d01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Println("AOC 2018 - day 01")

	startTime := time.Now()
	part1Result := Part1(inputPath)
	elapsed := time.Since(startTime)

	fmt.Printf("Part 1: %d\n%d μs\n", part1Result, elapsed.Microseconds())

	startTime = time.Now()
	part2Result := Part2(inputPath)
	elapsed = time.Since(startTime)

	fmt.Printf("Part 2: %d\n%d μs\n", part2Result, elapsed.Microseconds())
}

func Part1(inputPath string) (finalFrequency int) {
	file, err := os.Open(inputPath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		adjustment, err := parseLine(line)
		if err != nil {
			log.Fatalln(err)
		}
		finalFrequency += adjustment
	}
	return
}

func parseLine(line string) (change int, err error) {
	line = strings.TrimSpace(line)

	change, err = strconv.Atoi(line[1:])

	if line[0] == '-' && err == nil {
		change *= -1
	}
	return
}

func Part2(inputPath string) int {
	return 0
}

package y2024d11

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ddonukis/advent-of-code-golang/solutions/parsing"
)

func Solve(inputPath string) {
	data := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(data)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	data = parseInput(inputPath)

	t0 = time.Now()
	result2 := Part2(data)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func parseInput(inputPath string) []int {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	lines := parsing.ExtractInts(string(content))
	return lines[0]
}

func Part1(data []int) int {
	pebbles := data
	for i := 0; i < 25; i++ {
		fmt.Println("Starting iteration: ", i, "array len: ", len(pebbles))
		updated := make([]int, 0)
		for _, pebble := range pebbles {
			updated = append(updated, transformPebble(pebble)...)
		}
		pebbles = updated
	}
	return len(pebbles)
}

func splitCondition(pebble int) bool {
	pebStr := fmt.Sprintf("%d", pebble)
	return len(pebStr)%2 == 0
}

func transformPebble(pebble int) []int {
	if pebble == 0 {
		return []int{1}
	} else if splitCondition(pebble) {
		pebStr := fmt.Sprintf("%d", pebble)
		left, _ := strconv.Atoi(pebStr[:len(pebStr)/2])
		right, _ := strconv.Atoi(pebStr[len(pebStr)/2:])
		return []int{left, right}
	}
	return []int{pebble * 2024}
}

func Part2(data []int) int {
	countByPebble := make(map[int]int)
	for _, peb := range data {
		countByPebble[peb]++
	}

	for i := 0; i < 75; i++ {
		fmt.Println("Starting iteration: ", i, "map len: ", len(countByPebble))
		countByPebble2 := make(map[int]int)
		for peb := range countByPebble {
			initialCount := countByPebble[peb]
			resultingPebbles := transformPebble(peb)
			for _, peb := range resultingPebbles {
				countByPebble2[peb] += initialCount
			}
		}
		countByPebble = countByPebble2
	}

	sum := 0
	for peb := range countByPebble {
		sum += countByPebble[peb]
	}
	return sum
}

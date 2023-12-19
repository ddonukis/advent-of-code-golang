package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	checksum := 0
	part2Total := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := parseLine(line)

		slices.Sort(numbers)
		checksum += numbers[len(numbers)-1] - numbers[0]

		sum := part2Calc(numbers)
		part2Total += sum
	}
	fmt.Printf("Part 1: %d\n", checksum)
	fmt.Printf("Part 2: %d\n", part2Total)
}

func part2Calc(nums []int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		for j, n := range nums {
			if i == j {
				continue
			}
			if nums[i]%n == 0 {
				return nums[i] / n
			}
		}
	}
	return 0
}

func parseLine(line string) []int {
	cols := strings.Fields(line)
	nums := make([]int, len(cols))
	for i, n := range cols {
		val, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("bad nunber '%s': %v\n", n, err)
		}
		nums[i] = val
	}
	return nums
}

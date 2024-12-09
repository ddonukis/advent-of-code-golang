package y2024d09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Solve(inputPath string) {
	data := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(data)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(data)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func parseInput(inputPath string) []int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var nums []int
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		nums = make([]int, len(line))
		for i, dig := range line {
			n, err := strconv.Atoi(string(dig))
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}
		fmt.Printf("len(nums): %d\n", len(nums))
		return nums
	}
	return make([]int, 0)
}

func Part1(nums []int) int {
	mem := unfoldMemLayout(nums)

	lastIdx := len(mem) - 1
	for idx := 0; idx < len(mem); idx++ {
		block := mem[idx]
		if block.isFree {
			for i := lastIdx; i > idx; i-- {
				if mem[i].isFree {
					continue
				}
				mem[idx] = mem[i]
				mem[i] = block
				lastIdx = i
				break
			}
			if lastIdx <= idx {
				break
			}
		}
	}

	return checksum(mem)
}

func Part2(nums []int) int {
	mem := unfoldMemLayout(nums)
	return checksum(mem)
}

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

func findFreeSpan(mem []Block, startAt int, size int, stopAt int) int {
	curSize := 0
	idx := startAt
	for idx < stopAt {
		if mem[idx].isFree {
			curSize++
		} else {
			curSize = 0
		}
		if curSize == size {
			return idx - curSize + 1
		}
		idx++
	}
	return -1
}

func moveBlocks(startsAt int, size int8, mem []Block, leftmostFreeIdxBySize map[int8]int) (moved bool) {
	idx, found := leftmostFreeIdxBySize[size]

	if !found {
		// search from 0
	} else if found && idx > 0 && idx < startsAt && !mem[idx].isFree {
		// search from idx
	} else if found && (idx == -1 || idx > startsAt) {
		return false
		// no free space to the left of the file -> end
	} else {
		// can swap
	}

	for i := 0; i < int(size); i++ {
		mem[idx+i] = mem[startsAt+i]
		mem[startsAt+i] = Block{0, true}
	}

	return false
}

func Part2(nums []int) int {
	mem := unfoldMemLayout(nums)

	leftmostFreeIdxBySize := make(map[int8]int) // size -> leftmost idx

	var curFileSize int8 = 0
	var curFileId int16 = -1

	// 00...111...2...333.44.5555.6666.777.888899
	for i := len(mem) - 1; i > -1; i-- {
		block := mem[i]
		if block.isFree {
			if curFileId > 0 {
				fmt.Printf("id: %d, size: %d\n", curFileId, curFileSize)
				moveBlocks(i+1, curFileSize, mem, leftmostFreeIdxBySize)
				curFileId = -1
				curFileSize = 0
			}
		} else if block.id != curFileId {
			if curFileId > 0 {
				fmt.Printf("id: %d, size: %d\n", curFileId, curFileSize)
				moveBlocks(i+1, curFileSize, mem, leftmostFreeIdxBySize)
			}
			curFileId = block.id
			curFileSize = 1
		} else {
			curFileSize++
		}
	}
	if curFileId > -1 {
		fmt.Printf("id: %d, size: %d\n", curFileId, curFileSize)
	}

	// curFreeSize := 0
	// curFreeStartIdx := 0
	// for idx := 0; idx < len(mem); idx++ {
	// 	block := mem[idx]
	// 	if block.isFree {
	// 		if curFreeSize == 0 {
	// 			curFreeSize += 1
	// 			curFreeStartIdx = idx
	// 		}
	// 	} else {
	// 		leftmostFreeIdxBySize[curFreeSize] =
	// 	}

	// }

	return checksum(mem)
}

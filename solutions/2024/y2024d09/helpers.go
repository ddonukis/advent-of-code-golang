package y2024d09

import "fmt"

type Block struct {
	id     int16
	isFree bool
}

func (b Block) String() string {
	if b.isFree {
		return "."
	}
	return fmt.Sprintf("%d", b.id)
}

func unfoldMemLayout(nums []int) []Block {
	memLen := 0
	for _, blockCount := range nums {
		memLen += blockCount
	}
	mem := make([]Block, memLen)
	runningId := 0
	id := 0
	for idx, blockCount := range nums {
		isFree := idx%2 != 0
		for i := runningId; i < runningId+blockCount; i++ {
			if isFree {
				mem[i] = Block{0, isFree}
			}
			mem[i] = Block{int16(id), isFree}
		}
		if !isFree {
			id++
		}
		runningId += blockCount
	}
	return mem
}

func checksum(mem []Block) int {
	sum := 0
	for i, block := range mem {
		if !block.isFree {
			sum += i * int(block.id)
		}
	}
	return sum
}

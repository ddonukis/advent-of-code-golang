package y2024d02

import (
	"strconv"
	"strings"
)

func parseLine(line string) (numbers []int, err error) {
	rawNums := strings.Fields(line)
	numbers = make([]int, len(rawNums))
	for i, n := range rawNums {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			return numbers, err
		}
		numbers[i] = nInt
	}
	return
}

package parsing

import (
	"strconv"
	"strings"
	"unicode"
)

// For every line in text create a slice of ints extracted
func ExtractInts(text string) [][]int {
	lines := strings.SplitAfter(text, "\n")
	var numArrays [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		var numbers []int

		var curNumIdx, curNumLen int
		for i, ch := range line {
			if curNumLen == 0 && (unicode.IsDigit(ch) || ch == '-') {
				curNumIdx = i
				curNumLen++
			} else if curNumLen > 0 && unicode.IsDigit(ch) {
				curNumLen++
			} else if curNumLen > 0 {
				num, err := strconv.Atoi(line[curNumIdx : curNumIdx+curNumLen])
				if err == nil {
					numbers = append(numbers, num)
				}
				curNumLen = 0
			}
		}
		numArrays = append(numArrays, numbers)
	}
	return numArrays
}

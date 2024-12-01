package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	row   int
	col   int
	len   int
	value int
}

func (n *Number) hasAdjacentSymbol(symbolsMatrix *[][]bool) bool {
	leftmost := n.col
	if leftmost > 0 {
		leftmost -= 1
		// check left if it's not the first column
		if (*symbolsMatrix)[n.row][leftmost] {
			fmt.Printf("'%d' adjacent to [%d][%d] (left)\n", n.value, n.row, leftmost)
			return true
		}
	}

	rightmost := n.col + n.len
	if rightmost < len((*symbolsMatrix)[n.row]) {
		// check right if it's not the last column
		if (*symbolsMatrix)[n.row][rightmost] {
			fmt.Printf("'%d' adjacent to [%d][%d] (left)\n", n.value, n.row, rightmost)
			return true
		}
		rightmost += 1
	}

	fmt.Printf("leftmost: %d, rightmost: %d\n", leftmost, rightmost)
	// check above if it's not the first row
	if n.row > 0 {
		for col, test := range (*symbolsMatrix)[n.row-1][leftmost:rightmost] {
			if test {
				fmt.Printf("'%d' adjacent to [%d][%d] (above)\n", n.value, n.row-1, col)
				return true
			}
		}
	}
	// check below if it's not the last row
	if n.row+1 < len(*symbolsMatrix) {
		for col, test := range (*symbolsMatrix)[n.row+1][leftmost:rightmost] {
			if test {
				fmt.Printf("'%d' adjacent to [%d][%d] (below)\n", n.value, n.row+1, col)
				return true
			}
		}
	}
	fmt.Printf("'%d' not adjacent to any symbols\n", n.value)
	return false
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	linesProcessed := 0

	numbers := make([]Number, 0)
	symbols := make([][]bool, 0)

	sumPart1 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Printf("\n%s\n", line)

		parsedNumbers, symbolMask := parseLine(line, linesProcessed)
		numbers = append(numbers, parsedNumbers...)
		symbols = append(symbols, symbolMask)
		linesProcessed++
	}
	for _, n := range numbers {
		// fmt.Println(n)
		if n.hasAdjacentSymbol(&symbols) {
			sumPart1 += n.value
		}
	}
	fmt.Printf("Part 1: %d\n", sumPart1)

	Main2()
}

func parseLine(line string, lineNumber int) (numbers []Number, symbolMask []bool) {
	symbolMask = make([]bool, len(line))
	numbers = make([]Number, 0)
	numberStartIdx, number := 0, make([]rune, 0, len(line))

	for idx, char := range line {
		if unicode.IsDigit(char) {
			// accumulate digits
			if len(number) == 0 {
				numberStartIdx = idx
			}
			number = append(number, char)
		} else {
			// store number
			if len(number) > 0 {
				numVal, err := strconv.Atoi(string(number))
				if err != nil {
					fmt.Printf("Failed to convert %v to number: %v\n", number, err)
				}
				numbers = append(numbers, Number{row: lineNumber, col: numberStartIdx, len: len(number), value: numVal})
				// fmt.Prinln(Number{row: lineNumber, col: numberStartIdx, len: len(number), value: numVal})
				number = number[:0]
			}
			if char != '.' {
				symbolMask[idx] = true
			}
		}
	}
	if len(number) > 0 {
		numVal, err := strconv.Atoi(string(number))
		if err != nil {
			fmt.Printf("Failed to convert %v to number: %v\n", number, err)
		}
		numbers = append(numbers, Number{row: lineNumber, col: numberStartIdx, len: len(number), value: numVal})
		// fmt.Prinln(Number{row: lineNumber, col: numberStartIdx, len: len(number), value: numVal})
		number = number[:0]
	}
	return numbers, symbolMask
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func parseArgs() (filePath string, err error) {
	switch len(os.Args) {
	case 2:
		return os.Args[1], nil
	case 1:
		return "", errors.New("must proved file path as the program argument")
	default:
		return "", fmt.Errorf("only 1 argument expected, but too many arguments provided: %v", os.Args[1:])
	}
}

func main() {
	fp, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	fScanner := bufio.NewScanner(f)

	total := 0
	patternLinesBuffer := make([]string, 0, 10)
	for fScanner.Scan() {
		line := fScanner.Text()
		if len(line) == 0 {
			// reached end of pattern
			total += summarizePattern(patternLinesBuffer)
			patternLinesBuffer = patternLinesBuffer[:0]
			continue
		}
		patternLinesBuffer = append(patternLinesBuffer, line)
	}
	total += summarizePattern(patternLinesBuffer)
	fmt.Printf("Total: %d\n", total)
}

func summarizePattern(lines []string) int {
	// test horizontal symmatry
	sum := 0
	horizontalAxis := symmetryAxis(lines)
	if horizontalAxis >= 0 {
		fmt.Printf("Horizontal axis score: %d\n", (horizontalAxis+1)*100)
		sum += (horizontalAxis + 1) * 100
	}

	// test vertical symmetry
	cols := linesToCols(lines)
	verticalAxis := symmetryAxis(cols)
	if verticalAxis >= 0 {
		fmt.Printf("Vertical axis score: %d\n", verticalAxis+1)
		sum += verticalAxis + 1
	}

	return sum
}

func linesToCols(lines []string) []string {
	cols := make([]string, len(lines[0]))
	for colIdx := range cols {
		col := make([]rune, len(lines))
		for rowIdx := 0; rowIdx < len(lines); rowIdx++ {
			col[rowIdx] = []rune(lines[rowIdx])[colIdx]
		}
		cols[colIdx] = string(col)
	}
	return cols
}

// return index of a line after which the lines are mirrored,
// if no mirrored lines found return -1
func symmetryAxis(lines []string) int {
tryTestAxis:
	for testAxis := 0; testAxis+1 < len(lines); testAxis++ {
		for offset := 0; testAxis-offset >= 0 && testAxis+1+offset < len(lines); offset++ {
			// fmt.Printf("ta: %d, offset: %d\n", testAxis, offset)
			if lines[testAxis-offset] != lines[testAxis+1+offset] {
				continue tryTestAxis
			}
			// fmt.Printf("%s %s ta: %d, offset: %d\n", lines[testAxis-offset], lines[testAxis+1+offset], testAxis, offset)
		}
		return testAxis
	}
	return -1
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

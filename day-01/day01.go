package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC 2023, day 1!")
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputScanner := bufio.NewScanner(inputFile)
	var sumPart1, sumPart2 int = 0, 0
	for inputScanner.Scan() {

		line := inputScanner.Text()
		number, err := extractNumber(line)
		if err == nil {
			sumPart1 += number
		} else {
			log.Printf("Could not extract number from '%s': %v\n", line, err)
		}

		lineReplaced := replaceOverlappingWordsWithDigits(line)

		number, err = extractNumber(lineReplaced)
		if err == nil {
			sumPart2 += number
		} else {
			log.Printf("Could not extract number from '%s': %v\n", lineReplaced, err)
		}

	}
	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
}

type wordDigitLocation struct {
	value string
	start int
	end   int
}

// Replace digits spelled out as words with numeric digits. Accounts for overlapping words.
// E.g. "123fourfive" -> "12345"; "oneight" -> "18"; "sixteen" -> "6teen"
func replaceOverlappingWordsWithDigits(line string) string {
	digitWords := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	wordDigitBounds := make([]wordDigitLocation, 0)
	for i, w := range digitWords {
		for _, from := range substringIndices(line, w) {
			to := from + len(w)
			wordDigitBounds = append(wordDigitBounds, wordDigitLocation{strconv.Itoa(i + 1), from, to})
		}
	}
	if len(wordDigitBounds) == 0 {
		return line
	}
	slices.SortFunc(
		wordDigitBounds,
		func(a, b wordDigitLocation) int {
			return a.start - b.start
		},
	)
	replacedStringParts := make([]string, 0, len(line))
	var cursor int = 0
	for _, wd := range wordDigitBounds {
		if cursor < wd.start {
			replacedStringParts = append(replacedStringParts, line[cursor:wd.start])
		}
		replacedStringParts = append(replacedStringParts, wd.value)
		cursor = wd.end
	}
	if cursor < len(line) {
		replacedStringParts = append(replacedStringParts, line[cursor:])
	}
	replacedLine := strings.Join(replacedStringParts, "")
	return replacedLine
}

// Return a two digit number compsed from the first and last digit found in line.
// Note, if there's only one digit in the string it's both the first and the last one.
// E.g. "1bear6" -> 16; "blah1" -> 11
func extractNumber(line string) (int, error) {
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(line, -1)
	numberStr := digits[0] + digits[len(digits)-1]
	numuber, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}
	return numuber, nil
}

func substringIndices(s string, substr string) (indices []int) {
	indices = make([]int, 0, len(s))
	cursor := 0
	for {
		idx := strings.Index(s[cursor:], substr)
		if idx < 0 {
			break
		}
		sIdx := cursor + idx
		indices = append(indices, sIdx)
		cursor = sIdx + len(substr)
	}
	return
}

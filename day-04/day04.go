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

var cardIdRe = regexp.MustCompile(`^\s*Card\s+(\d+):`)

func main() {
	const filename = "data.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fScanner := bufio.NewScanner(file)

	totalPoints := 0
	for fScanner.Scan() {
		line := fScanner.Text()
		cardNo, count := countWinningNumbers(line)
		points := calcPoints(count)
		fmt.Printf("Card %d: %d\n", cardNo, points)
		totalPoints += points
	}
	fmt.Printf("\nTotal (Part 1): %d\n", totalPoints)
}

func countWinningNumbers(line string) (cardId, count int) {
	matches := cardIdRe.FindStringSubmatch(line)
	if len(matches) == 2 {
		cardId, _ = strconv.Atoi(matches[1])
	} else {
		log.Printf("Invalid line: '%s'\n", line)
		return 0, 0
	}
	line = line[len(matches[0]):] // drop the card number prefix
	lineElems := strings.Fields(strings.TrimSpace(line))
	separatorIdx := slices.Index(lineElems, "|")
	if separatorIdx == -1 {
		log.Printf("No separator found in line: '%s'\n", line)
		return 0, 0
	}

	winningNumbers := make([]int, separatorIdx)
	for i, winningNumStr := range lineElems[:separatorIdx] {
		n, err := strconv.Atoi(winningNumStr)
		if err != nil {
			log.Printf("Could not parse winning number: '%s'\n", winningNumStr)
			return 0, 0
		}
		winningNumbers[i] = n
	}

	matchingNumbers := 0
	for _, tryNum := range lineElems[separatorIdx+1:] {
		n, err := strconv.Atoi(tryNum)
		if err != nil {
			log.Printf("Could not parse try number: '%s'\n", tryNum)
			return 0, 0
		}
		if slices.Index(winningNumbers, n) != -1 {
			matchingNumbers++
		}
	}
	return cardId, matchingNumbers
}

func calcPoints(matchCount int) int {
	if matchCount == 0 {
		return 0
	}
	points := 1
	for i := 1; i < matchCount; i++ {
		points *= 2
	}
	return points
}

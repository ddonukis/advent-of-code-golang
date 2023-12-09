package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFilePath() (string, error) {
	switch len(os.Args[1:]) {
	case 1:
		return os.Args[1], nil
	case 0:
		return "", errors.New("no arguments passed")
	default:
		return "", errors.New("too many arguments, 1 expected")
	}
}

func iterParsedLines[T any](filepath string, parseFunc func(string) T) <-chan T {
	lines := make(chan T)
	go func() {
		defer close(lines)
		f, err := os.Open(filepath)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			lines <- parseFunc(fs.Text())
		}
	}()
	return lines
}

func main() {
	path, err := getFilePath()
	if err != nil {
		log.Fatalf("Cannot get input file path: %v\n", err)
	}
	fmt.Printf("Input: '%s'\n", path)

	total := 0
	totalPart2 := 0
	for numbers := range iterParsedLines(path, toInts) {
		total += predictNextNum(numbers)
		totalPart2 += predictPrevNum(numbers)
	}
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", totalPart2)
}

func predictNextNum(numbers []int) int {
	// fmt.Println("Predicting...")
	// fmt.Println(numbers)

	var diffed []int = numbers
	allZero := false
	predicted := numbers[len(numbers)-1]
	for !allZero {
		diffed, allZero = diffNums(diffed)
		predicted += diffed[len(diffed)-1]
		// fmt.Println(diffed)
		// fmt.Printf("Predicted: %d\n", predicted)
	}
	return predicted
}

func predictPrevNum(numbers []int) int {
	// fmt.Println("\nPredicting...")
	// fmt.Println(numbers)

	var diffed []int = numbers
	allZero := false
	firstNums := make([]int, 1, len(numbers))

	firstNums[0] = numbers[0]
	for !allZero {
		diffed, allZero = diffNums(diffed)
		firstNums = append(firstNums, diffed[0])
		// fmt.Println(diffed)
		// fmt.Printf("first nums: %v\n", firstNums)
	}

	for i := len(firstNums) - 2; i > 0; i-- {
		firstNums[i-1] -= firstNums[i]
	}

	// fmt.Printf("%v -> %d\n", numbers, firstNums[0])
	return firstNums[0]
}

func diffNums(numbers []int) (diffed []int, allZero bool) {
	diffed = make([]int, len(numbers)-1)
	allZero = true
	for i := 1; i < len(numbers); i++ {
		diffed[i-1] = numbers[i] - numbers[i-1]
		if diffed[i-1] != 0 {
			allZero = false
		}
	}
	return diffed, allZero
}

func toInts(line string) []int {
	strInts := strings.Fields(line)
	ints := make([]int, len(strInts))
	for i, val := range strInts {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("'%s' is not a valid number: %v\n", val, err)
		}
		ints[i] = intVal
	}
	return ints
}

package y2024d01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Println("AOC 2024 - day 01")

	startTime := time.Now()
	part1Result := Part1(inputPath)
	elapsed := time.Since(startTime)

	fmt.Printf("Part 1: %d\n%d μs\n", part1Result, elapsed.Microseconds())

	startTime = time.Now()
	part2Result := Part2(inputPath)
	elapsed = time.Since(startTime)

	fmt.Printf("Part 2: %d\n%d μs\n", part2Result, elapsed.Microseconds())
}

func Part1(dataFilePath string) int {
	fmt.Println("AOC 2024 - day 01")

	file, err := os.Open(dataFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numsA := make([]int, 0)
	numsB := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		a, b, err := parseLine(line)
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Printf("%d, %d\n", a, b)
		numsA = append(numsA, a)
		numsB = append(numsB, b)
	}
	sort.Ints(numsA)
	sort.Ints(numsB)

	sum := 0
	for i, a := range numsA {
		b := numsB[i]

		diff := a - b
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return sum
}

func Part2(dataFilePath string) int {
	fmt.Println("AOC 2024 - day 01")

	file, err := os.Open(dataFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numsA := make([]int, 0)
	numsB := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		a, b, err := parseLine(line)
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Printf("%d, %d\n", a, b)
		numsA = append(numsA, a)
		numsB[b] += 1
		// fmt.Printf("b: %d\n%v\n", b, numsB)
	}

	sum := 0
	for _, a := range numsA {
		appearances := numsB[a]

		score := a * appearances
		// fmt.Printf("%d * %d = %d\n", a, appearances, score)
		sum += score
	}

	return sum
}

func parseLine(line string) (a, b int, err error) {
	line = strings.TrimSpace(line)
	first_num, second_num, found := strings.Cut(line, " ")
	second_num = strings.TrimSpace(second_num)

	if !found {
		err = fmt.Errorf("expected two whitespace separted integers, got: %s", line)
	}
	var first_err, second_err error
	a, first_err = strconv.Atoi(first_num)
	b, second_err = strconv.Atoi(second_num)
	if first_err != nil || second_err != nil {
		err = fmt.Errorf("could not convert '%s' and/or '%s' to intrgers", first_num, second_num)
		return
	}
	return
}

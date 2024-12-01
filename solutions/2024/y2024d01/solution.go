package y2024d01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(dataFilePath string) {
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

	fmt.Printf("Part 1: %d\n", sum)
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	const path = "input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	seeds, almanac := extractData(scanner)

	fmt.Println(seeds)
	printMap(almanac)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := traceSeedToLocation(seed, almanac)
		fmt.Printf("seed %d -> location %d\n", seed, location)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Printf("Lowest location (Part 1): %d\n", lowestLocation)
}

type Range struct {
	source      int
	destination int
	length      int
}

func (r Range) String() string {
	return fmt.Sprintf("Range{source: %d, destination: %d, length: %d}", r.source, r.destination, r.length)
}

func extractData(scanner *bufio.Scanner) (seeds []int, almanac map[string][]Range) {
	const seedsPrefix = "seeds:"
	const mapSuffix = "map:"
	seeds = make([]int, 0, 10)
	readingMap := false
	var currentHeader string
	currentMap := make([]Range, 0, 5)
	almanac = make(map[string][]Range)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		isSeedsList := strings.HasPrefix(trimmedLine, seedsPrefix)
		if isSeedsList {
			for _, numStr := range strings.Fields(trimmedLine[len(seedsPrefix):]) {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Printf("Cannot convert '%s' to int: %v\n", numStr, err)
					continue
				}
				seeds = append(seeds, num)
			}
		}

		isMapHeader := strings.HasSuffix(trimmedLine, mapSuffix)
		if isMapHeader {
			currentHeader = trimmedLine[:len(trimmedLine)-len(mapSuffix)-1]
			readingMap = true
			continue
		}

		// maps are separated by empty spaces
		if readingMap && len(trimmedLine) == 0 {
			slices.SortFunc(currentMap, func(a, b Range) int { return a.source - b.source })
			almanac[currentHeader] = currentMap
			currentMap = make([]Range, 0, 5)
			readingMap = false
			continue
		}

		if readingMap {
			mapRowStr := strings.Fields(trimmedLine)
			if len(mapRowStr) != 3 {
				fmt.Printf("Invalid map row: '%s'\n", trimmedLine)
				continue
			}
			mapValues := make([]int, len(mapRowStr))
			for idx, elem := range mapRowStr {
				val, err := strconv.Atoi(elem)
				if err != nil {
					fmt.Printf("Cannot convert '%s' to int: %v\n", elem, err)
					continue
				}
				mapValues[idx] = val
			}
			rowRange := Range{source: mapValues[1], destination: mapValues[0], length: mapValues[2]}
			currentMap = append(currentMap, rowRange)
		}
	}
	if readingMap {
		almanac[currentHeader] = currentMap
	}
	return
}

func printMap[T any](m map[string]T) {
	for key, value := range m {
		fmt.Printf("'%s':\n  %v\n", key, value)
	}
}

func findNextMap(key string, almanac map[string][]Range) (nextKey string, found bool) {
	fromTo := strings.Split(key, "-to-")

	nextPrefix := fromTo[len(fromTo)-1]

	for k := range almanac {
		if strings.HasPrefix(k, nextPrefix) {
			return k, true
		}
	}
	return "", false
}

func traceSeedToLocation(seed int, almanac map[string][]Range) (location int) {
	nextId := seed
	nextKey := "seed"
	var found bool
	for {
		nextKey, found = findNextMap(nextKey, almanac)
		if found {
			nextId = getDestination(nextId, almanac[nextKey])
		} else {
			location = nextId
			break
		}
	}
	return location
}

func getDestination(source int, mapping []Range) (destination int) {
	for _, r := range mapping {
		if r.source <= source && source < r.source+r.length {
			return r.destination + source - r.source
		}
	}
	return source
}

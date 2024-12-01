package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

func (cubeSet CubeSet) isPossible(maxCubeSet CubeSet) bool {
	if cubeSet.red > maxCubeSet.red {
		return false
	}
	if cubeSet.green > maxCubeSet.green {
		return false
	}
	if cubeSet.blue > maxCubeSet.blue {
		return false
	}
	return true
}

func (cubeSet CubeSet) String() string {
	return fmt.Sprintf("{R=%d, G=%d, B=%d}", cubeSet.red, cubeSet.green, cubeSet.blue)
}

func NewCubeSetFromString(s string) (cubeSet CubeSet, err error) {
	cubeRe := regexp.MustCompile(`\s*(\d+)\s+([a-z]+)`)
	for _, cubeMatch := range cubeRe.FindAllStringSubmatch(s, -1) {
		if len(cubeMatch) < 3 {
			return cubeSet, fmt.Errorf("invalid string format: '%s'", s)
		}
		count, _ := strconv.Atoi(cubeMatch[1])
		switch cubeMatch[2] {
		case "red":
			cubeSet.red += count
		case "green":
			cubeSet.green += count
		case "blue":
			cubeSet.blue += count
		default:
			return cubeSet, fmt.Errorf("invalid color: '%s'", cubeMatch[2])
		}
	}
	return
}

type Game struct {
	id   int
	sets []CubeSet
}

func (game *Game) isPossible(maxCubeSet CubeSet) bool {
	for _, set := range game.sets {
		if !set.isPossible(maxCubeSet) {
			return false
		}
	}
	return true
}

func (game *Game) setPower() int {
	var maxCubeSet CubeSet
	for _, set := range game.sets {
		if set.red > maxCubeSet.red {
			maxCubeSet.red = set.red
		}
		if set.green > maxCubeSet.green {
			maxCubeSet.green = set.green
		}
		if set.blue > maxCubeSet.blue {
			maxCubeSet.blue = set.blue
		}
	}
	return maxCubeSet.red * maxCubeSet.green * maxCubeSet.blue
}

func (game *Game) String() string {
	return fmt.Sprintf("<id=%d, sets=%v>", game.id, game.sets)
}

func NewGameFromLine(line string) (*Game, error) {
	var game Game
	var err error
	prefixRe := regexp.MustCompile(`^Game\s+(\d+):\s+`)
	matches := prefixRe.FindStringSubmatch(line)
	if len(matches) < 2 {
		return nil, fmt.Errorf("invalid line format: '%s'", line)
	}
	game.id, err = strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("Could not convert game id to number: %v", err)
	}
	line, _ = strings.CutPrefix(line, matches[0])
	cubeSetStrings := strings.Split(line, ";")
	game.sets = make([]CubeSet, len(cubeSetStrings))
	for i, cs := range cubeSetStrings {
		cubeSet, _ := NewCubeSetFromString(cs)
		game.sets[i] = cubeSet
	}
	return &game, nil
}

func main() {
	maxCubeSet := CubeSet{red: 12, green: 13, blue: 14}

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Could not open the file: %v", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	part1sum := 0
	part2sum := 0
	for fileScanner.Scan() {

		line := fileScanner.Text()
		game, err := NewGameFromLine(line)
		if err != nil {
			log.Fatalf("Could not parse game from line: %v", err)
		}
		fmt.Printf("line: '%s'\n  -> %v\n", line, game)
		if game.isPossible(maxCubeSet) {
			part1sum += game.id
		}
		part2sum += game.setPower()

	}
	fmt.Printf("\nPart 1: %d\n", part1sum)
	fmt.Printf("Part 2: %d\n", part2sum)

}

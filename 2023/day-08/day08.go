package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type Fork struct {
	left  string
	right string
}

func main() {
	const filepath = "input.txt"

	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`([0-9A-Z]{3})`)

	var commands string
	nodes := make(map[string]Fork)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		if len(commands) == 0 {
			commands = line
			continue
		}
		if len(line) == 0 {
			continue
		}
		node := re.FindAllString(line, 3)

		if len(node) != 3 {
			log.Fatalf("Invalid line '%s'\n", line)
		}
		nodes[node[0]] = Fork{left: node[1], right: node[2]}
	}
	// part1(commands, nodes)
	part2(commands, nodes)
}

func part1(commands string, nodes map[string]Fork) {
	nextNode := "AAA"
	var stepsTaken, totalSteps int

	startTime := time.Now()
	fmt.Printf("Started at %v\n", startTime)

	for len(nextNode) > 0 {
		stepsTaken, nextNode = walkNodes(nextNode, commands, nodes)
		totalSteps += stepsTaken
	}
	elapsed := time.Since(startTime)
	fmt.Printf("Done at %v (%f seconds)\n", time.Now(), elapsed.Seconds())
	fmt.Printf("Steps took: %d\n", totalSteps)
}

func part2(commands string, nodes map[string]Fork) {
	startTime := time.Now()
	fmt.Printf("Part 2 started at %v\n", startTime)

	cycleLengths := make([]int, 0)
	for nodeKey := range nodes {
		if nodeKey[2] == 'A' {
			var steps, cycleSteps int
			var nextNode string = nodeKey
			for len(nextNode) > 0 {
				steps, nextNode = walkNodesP2(nextNode, commands, nodes)
				cycleSteps += steps
			}
			fmt.Printf("%s cycle %d\n", nodeKey, cycleSteps)
			cycleLengths = append(cycleLengths, cycleSteps)
		}
	}
	fmt.Println(cycleLengths)

	fmt.Println(lcmSlice(cycleLengths))

	elapsed := time.Since(startTime)
	fmt.Printf("Part 2 done at %v (%f seconds)\n", time.Now(), elapsed.Seconds())
	// fmt.Printf("Steps took: %d\n", totalSteps)
}

func walkNodes(startAt string, commands string, nodes map[string]Fork) (steps int, nextKey string) {
	nextKey = startAt
	var command rune
	for steps, command = range commands {
		// fmt.Printf("% 6d: '%s' -%c-> ", steps+1, nextKey, command)
		node, found := nodes[nextKey]
		if !found {
			log.Fatalf("Node '%v' not found\n", nextKey)
		}
		switch command {
		case 'L':
			nextKey = node.left
		case 'R':
			nextKey = node.right
		default:
			log.Fatalf("Invalid command '%v'\n", command)
		}
		// fmt.Printf("'%s'\n", nextKey)
		if nextKey == "ZZZ" {
			return 1 + steps, ""
		}
	}
	return steps + 1, nextKey
}

func walkNodesP2(startAt string, commands string, nodes map[string]Fork) (steps int, nextKey string) {
	nextKey = startAt
	var command rune
	for steps, command = range commands {
		// fmt.Printf("% 6d: '%s' -%c-> ", steps+1, nextKey, command)
		node, found := nodes[nextKey]
		if !found {
			log.Fatalf("Node '%v' not found\n", nextKey)
		}
		switch command {
		case 'L':
			nextKey = node.left
		case 'R':
			nextKey = node.right
		default:
			log.Fatalf("Invalid command '%v'\n", command)
		}
		// fmt.Printf("'%s'\n", nextKey)
		if nextKey[2] == 'Z' {
			return 1 + steps, ""
		}
	}
	return steps + 1, nextKey
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func lcmSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcm(result, number)
	}
	return result
}

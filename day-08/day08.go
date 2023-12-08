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

	re := regexp.MustCompile(`([A-Z]{3})`)

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

	// fmt.Printf("%d\n%v\n", len(nodes), nodes)
	// fmt.Printf("Start at '%s'\n", firstNodeKey)

	var stepsTaken, totalSteps int
	var nextNode string = "AAA"

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

func walkNodes(startAt string, commands string, nodes map[string]Fork) (steps int, nextKey string) {
	nextKey = startAt
	var command rune
	for steps, command = range commands {
		fmt.Printf("% 6d: '%s' -%c-> ", steps+1, nextKey, command)
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
		fmt.Printf("'%s'\n", nextKey)
		if nextKey == "ZZZ" {
			return 1 + steps, ""
		}
	}
	return steps + 1, nextKey
}

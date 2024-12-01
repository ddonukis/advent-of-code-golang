package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rootCandidates := make(map[string]bool)
	children := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			continue
		}
		parentName := strings.Fields(parts[0])[0]
		rootCandidates[parentName] = true
		children[parentName] = strings.Split(strings.TrimSpace(parts[1]), ", ")

	}

	for _, childrenNodes := range children {
		for _, childNode := range childrenNodes {
			isCand, found := rootCandidates[childNode]
			if found && isCand {
				rootCandidates[childNode] = false
			}
		}
	}

	for nodeName, isCandidate := range rootCandidates {
		if !isCandidate {
			continue
		}
		fmt.Printf("Root: %s\n", nodeName)
	}
}

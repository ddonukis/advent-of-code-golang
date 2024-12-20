package y2024d19

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Solve(inputPath string) {
	data := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(data)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(data)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func Part1(data InputData) int {
	slices.Sort(data.patterns)

	cleanedPatterns := make([]string, 0)
	for i, pat := range data.patterns {
		// fmt.Printf("pat: %q\n", pat)
		c := make([]string, 0, len(data.patterns)-1)
		c = append(c, data.patterns[:i]...)
		c = append(c, data.patterns[i+1:]...)
		// fmt.Printf("len(c): %d (%d)\n", len(c), len(data.patterns))
		// fmt.Printf("%q\n", c)
		if !matchLine(pat, c) {
			cleanedPatterns = append(cleanedPatterns, pat)
		} else {
			// fmt.Printf("redundant: %q\n", pat)
		}
	}

	fmt.Printf("patter count: %d -> %d\n", len(data.patterns), len(cleanedPatterns))

	sum := 0
	for _, line := range data.lines {
		// if line != "burrwrgrbuwwwrwbuurrbwgbgwburrwwugubuggbguwrgurbggurwrw" {
		// 	continue
		// }
		fmt.Printf("%q\n", line)
		if matchLine(line, cleanedPatterns) {
			fmt.Println("  -> ok")
			sum++
		} else {
			fmt.Println("  -> no match")
		}
	}

	return sum
}

type matchState struct {
	lineFromIdx int
	candidates  []string
}

func matchLine(line string, patterns []string) (found bool) {
	// fmt.Printf("Line: %q\n", line)
	startIdx := 0
	states := make([]matchState, 0)
	candidates := findCandidates(line[startIdx:], patterns)
	// fmt.Printf("first cands: %q\n", candidates)
	if len(candidates) == 0 {
		return false
	}
	states = append(states, matchState{startIdx, candidates})
	var curStateIdx int
	var curLine string
	var curPat string
	i := 0
	for len(states) > 0 {
		// for len(states) > 0 && i < 50 {
		i++
		curStateIdx = len(states) - 1
		// fmt.Printf("\nlen(states): %d; curState: %v\n", len(states), states[curStateIdx])

		// fmt.Printf("candidates: %q\n", states[curStateIdx].candidates)

		if len(states[curStateIdx].candidates) == 0 {
			states = slices.Delete(states, len(states)-1, len(states))
			// fmt.Printf("No candidates, removing cur state. Len sates: %d\n", len(states))
			continue
		}
		curPat = Pop(&states[curStateIdx].candidates)
		// fmt.Printf("curPat: %q, candidates: %q\n", curPat, states[curStateIdx].candidates)
		curLine = line[states[curStateIdx].lineFromIdx+len(curPat):]
		if len(curLine) == 0 {
			return true
		}
		cands := findCandidates(curLine, patterns)
		if len(cands) > 0 {
			// fmt.Printf("found cands: %q\n", cands)
			st := matchState{states[curStateIdx].lineFromIdx + len(curPat), cands}
			// fmt.Printf("new state: %v\n", st)
			states = append(states, st)
		} else {
			// fmt.Printf("did not find candidates, cands: %q\n", states[curStateIdx].candidates)
		}
	}
	return false
}

func Pop[T any](s *[]T) T {
	n := len(*s)
	item := (*s)[n-1]

	*s = slices.Delete(*s, n-1, n)
	return item
}

func findCandidates(line string, patterns []string) []string {
	candidates := make([]string, 0)
	for _, pat := range patterns {
		if strings.HasPrefix(line, pat) {
			candidates = append(candidates, pat)
		}
	}
	return candidates
}

func Part2(data InputData) int {
	return 0
}

func parseInput(inputPath string) InputData {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	patternsSection := true
	var patterns []string
	var lines []string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		line := scanner.Text()
		if len(line) == 0 {
			patternsSection = false
			continue
		}

		if patternsSection {
			patterns = strings.Split(line, ", ")
		} else {
			lines = append(lines, line)
		}
	}
	return InputData{patterns: patterns, lines: lines}
}

type InputData struct {
	patterns []string
	lines    []string
}

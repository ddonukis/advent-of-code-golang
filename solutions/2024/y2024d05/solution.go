package y2024d05

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Printf("AoC 2024 - day 5\n\n")

	t0 := time.Now()
	result1 := Part1(inputPath)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(inputPath)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func Part1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	isRules := true

	ruleBook := NewRuleBook()
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if isRules && line == "" {
			isRules = false
			continue
		}

		if isRules {
			rule, err := parseRule(line)
			if err != nil {
				log.Fatalln(err)
			}
			ruleBook.AddRule(rule)
		} else {
			pageNums, err := parsePageNums(line)
			if err != nil {
				log.Fatalln(err)
			}
			isValid := ruleBook.IsValidOrder(pageNums)
			// fmt.Printf("Line '%s' is %t\n", line, isValid)
			if isValid {
				sum += pageNums[len(pageNums)/2]
			}
		}
	}
	return sum
}

type Rule struct {
	before int
	after  int
}

func (r *Rule) IsSatisfied(pageNums []int) bool {
	beforeIdx := slices.Index(pageNums, r.before)
	if beforeIdx == -1 {
		return true
	}
	afterIdx := slices.Index(pageNums, r.after)
	if afterIdx == -1 {
		return true
	}
	return beforeIdx < afterIdx
}

type RuleBook struct {
	Index map[int][]Rule
}

func NewRuleBook() *RuleBook {
	idx := make(map[int][]Rule)
	return &RuleBook{idx}
}

func (ruleBook *RuleBook) AddRule(rule Rule) {
	ruleBook.Index[rule.before] = append(ruleBook.Index[rule.before], rule)
	ruleBook.Index[rule.after] = append(ruleBook.Index[rule.after], rule)
}

func (ruleBook *RuleBook) IsValidOrder(pageNums []int) bool {
	for _, pageNum := range pageNums {
		rules, found := ruleBook.Index[pageNum]
		if !found {
			continue
		}

		for _, rule := range rules {
			if !rule.IsSatisfied(pageNums) {
				return false
			}

		}
	}
	return true
}

func (ruleBook *RuleBook) ReorderPages(pageNums []int) []int {
	sortedNums := slices.Clone(pageNums)

	slices.SortFunc(sortedNums, func(a, b int) int {
		rules, exists := ruleBook.Index[a]
		if exists {
			for _, rule := range rules {
				if rule.before == a && rule.after == b {
					return -1
				}
				if rule.before == b && rule.after == a {
					return 1
				}
			}
		}
		return 0
	})
	return sortedNums
}

func parseRule(line string) (rule Rule, err error) {
	left, right, ok := strings.Cut(line, "|")
	if !ok {
		return Rule{}, errors.New("could not split the rule string")
	}
	lNum, err := strconv.Atoi(left)
	if err != nil {
		return Rule{}, err
	}
	rNum, err := strconv.Atoi(right)
	if err != nil {
		return Rule{}, err
	}
	return Rule{lNum, rNum}, nil

}

func parsePageNums(line string) (pageNums []int, err error) {
	rawNums := strings.Split(line, ",")
	pageNums = make([]int, len(rawNums))
	for i, rawN := range rawNums {
		n, err := strconv.Atoi(rawN)
		if err != nil {
			return pageNums, err
		}
		pageNums[i] = n
	}
	return
}

func Part2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	isRules := true

	ruleBook := NewRuleBook()
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if isRules && line == "" {
			isRules = false
			continue
		}

		if isRules {
			rule, err := parseRule(line)
			if err != nil {
				log.Fatalln(err)
			}
			ruleBook.AddRule(rule)
		} else {
			pageNums, err := parsePageNums(line)
			if err != nil {
				log.Fatalln(err)
			}
			reorderedPageNums := ruleBook.ReorderPages(pageNums)
			if !slices.Equal(pageNums, reorderedPageNums) {
				sum += reorderedPageNums[len(pageNums)/2]
			}
		}
	}
	return sum
}

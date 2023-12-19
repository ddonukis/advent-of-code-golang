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

var workflowPattern = regexp.MustCompile(`^([a-z]+)\{([^}]+)\}$`)
var partAttributePattern = regexp.MustCompile(`[amsx]=(\d+)`)

func main() {
	fp := os.Args[1]
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	workflows := make(map[string]Workflow)
	parts := make([]Part, 0, 10)

	for s.Scan() {
		line := s.Text()

		if len(line) == 0 {
			continue
		}
		if line[0] == '{' {
			part := parsePart(line)
			parts = append(parts, part)
		} else {
			wfName, wf := parseWorkflow(line)
			workflows[wfName] = wf
		}
	}

	total := 0
	for _, part := range parts {
		if checkPart(part, workflows) {
			total += part.Rating()
		}
	}
	fmt.Printf("Part 1: %d\n", total)
}

func parsePart(line string) Part {
	matches := partAttributePattern.FindAllStringSubmatch(line, -1)
	if len(matches) != 4 || len(matches[0]) != 2 {
		log.Fatalf("invalid part '%s'\n", line)
	}
	attrs := make([]int, len(matches))
	for i, m := range matches {
		val, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalf("invalid part '%s': %v\n", line, err)
		}
		attrs[i] = val
	}
	return Part{
		x: attrs[0],
		m: attrs[1],
		a: attrs[2],
		s: attrs[3],
	}
}

func parseWorkflow(line string) (name string, rules []Rule) {
	matches := workflowPattern.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 || len(matches[0]) != 3 {
		log.Fatalf("Invalid workflow '%s'\n", line)
	}
	name = matches[0][1]
	rulesRaw := strings.Split(matches[0][2], ",")
	rules = make([]Rule, len(rulesRaw))
	for i := range rules {
		rules[i] = parseRule(rulesRaw[i])
	}
	return
}

func parseRule(rule string) Rule {
	parts := strings.Split(rule, ":")
	if len(parts) == 1 {
		return &AlwaysTrue{outcome: parts[0]}
	} else {
		val, err := strconv.Atoi(parts[0][2:])
		if err != nil {
			log.Fatalf("cannot parse rule '%s': %v\n", rule, err)
		}
		switch parts[0][1] {
		case '>':
			return &GreaterThan{attribute: rune(parts[0][0]), value: val, outcome: parts[len(parts)-1]}
		case '<':
			return &LessThan{attribute: rune(parts[0][0]), value: val, outcome: parts[len(parts)-1]}
		}
	}
	return nil
}

type Rule interface {
	Check(Part) (bool, string)
}

type GreaterThan struct {
	attribute rune
	value     int
	outcome   string
}

func (rule *GreaterThan) Check(part Part) (matches bool, outcome string) {
	if part.LookupAttribute(rule.attribute) > rule.value {
		return true, rule.outcome
	}
	return false, ""
}

type LessThan struct {
	attribute rune
	value     int
	outcome   string
}

func (rule *LessThan) String() string {
	return fmt.Sprintf("Rule(%c < %d ? %s)", rule.attribute, rule.value, rule.outcome)
}
func (rule *GreaterThan) String() string {
	return fmt.Sprintf("Rule(%c > %d ? %s)", rule.attribute, rule.value, rule.outcome)
}
func (rule *AlwaysTrue) String() string {
	return fmt.Sprintf("Rule(true ? %s)", rule.outcome)
}

func (rule *LessThan) Check(part Part) (matches bool, outcome string) {
	if part.LookupAttribute(rule.attribute) < rule.value {
		return true, rule.outcome
	}
	return false, ""
}

type AlwaysTrue struct {
	outcome string
}

func (rule *AlwaysTrue) Check(part Part) (matches bool, outcome string) {
	return true, rule.outcome
}

type Workflow []Rule

type Part struct {
	x, m, a, s int
}

func (part *Part) Rating() int {
	return part.x + part.m + part.a + part.s
}

func (part *Part) LookupAttribute(name rune) int {
	switch name {
	case 'x':
		return part.x
	case 'm':
		return part.m
	case 'a':
		return part.a
	case 's':
		return part.s
	default:
		log.Fatalf("Invalid attribute '%c'\n", name)
	}
	return 0
}

func checkPart(part Part, workflows map[string]Workflow) (accepted bool) {
	nextWorkflowKey := "in"

iterateWorkflows:
	for !(nextWorkflowKey == "A" || nextWorkflowKey == "R") {
		wf := workflows[nextWorkflowKey]
		for _, rule := range wf {
			matched, nextKey := rule.Check(part)
			if matched {
				nextWorkflowKey = nextKey
				continue iterateWorkflows
			}
		}
	}
	return nextWorkflowKey == "A"
}

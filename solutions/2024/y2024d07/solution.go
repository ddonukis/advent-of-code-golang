package y2024d07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
	total   int
	numbers []int
}

func (eq Equation) TrySolve() (operators []Operator, isPossible bool) {
	operators = make([]Operator, len(eq.numbers)-1)

	fmt.Printf("%d = %v\n", eq.total, eq.numbers)
	for i := range operators {
		operators[i] = OP_ADD
	}

	for {
		res := eq.numbers[0]
		for i, op := range operators {
			res = op.Apply(res, eq.numbers[i+1])
			if res > eq.total {
				break
			}
		}
		// fmt.Printf("expected: %d, got: %d\n", eq.total, res)
		if res == eq.total {
			isPossible = true
			return
		}
		if !NextOperatorPermutation(operators) {
			break
		}
	}

	return operators, false
}

func NextOperatorPermutation(ops []Operator) (isNextAvailable bool) {
	if len(ops) == 0 {
		return false
	}

	for i := len(ops) - 1; i >= 0; i-- {
		if !ops[i].IsLast() {
			ops[i] = ops[i].Next()
			return true
		}
		ops[i] = ops[i].Next()
	}
	return false
}

type Operator uint8

const (
	OP_ADD Operator = iota
	OP_MUL
)

func (op Operator) Next() (nextOp Operator) {
	switch op {
	case OP_ADD:
		return OP_MUL
	case OP_MUL:
		return OP_ADD
	default:
		panic("invalid operator")
	}
}

func (op Operator) IsLast() bool {
	return op == OP_MUL
}

func (op Operator) Apply(a, b int) int {
	switch op {
	case OP_ADD:
		return a + b
	case OP_MUL:
		return a * b
	}
	return 0
}

func Solve(inputPath string) {
	equations := parseInput(inputPath)

	// fmt.Printf("%v\n", equations)

	t0 := time.Now()
	result1 := Part1(equations)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(equations)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func parseInput(inputPath string) []Equation {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	equations := make([]Equation, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		eq := parseEquation(line)
		equations = append(equations, eq)
	}

	return equations
}

func parseEquation(line string) Equation {
	total, numbers, found := strings.Cut(line, ":")
	if !found {
		panic("invalid line format %q")
	}
	t, err := strconv.Atoi(total)
	if err != nil {
		panic(err)
	}
	splitNums := strings.Fields(numbers)
	nums := make([]int, len(splitNums))
	for i, rawN := range splitNums {
		n, err := strconv.Atoi(rawN)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}

	return Equation{total: t, numbers: nums}
}

func Part1(equations []Equation) (sum int) {
	for _, equation := range equations {
		_, isPossible := equation.TrySolve()

		if isPossible {
			sum += equation.total
		}
	}
	return sum
}

func Part2(equations []Equation) int {
	return 0
}

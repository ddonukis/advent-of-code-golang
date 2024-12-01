package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type GearPostion struct {
	row int
	col int
}

func (g GearPostion) String() string {
	return fmt.Sprintf("<%d, %d>", g.row, g.col)
}

// seach around the gear and find all numbers
// if we have two different numbers multiply them and return as ratio
func (g GearPostion) Ratio(numbersMatrix *[][]*int) (ratio int, exists bool) {
	surroundingNums := make([]*int, 0)
	ownRow := (*numbersMatrix)[g.row]
	// check cell to the left if exists
	leftmost := g.col
	if leftmost > 0 {
		leftmost -= 1
		if left := ownRow[leftmost]; left != nil {
			surroundingNums = appendUnique(surroundingNums, left)
		}
	}
	// check cell to the right if exists
	rightmost := g.col
	if rightmost+1 < len(ownRow) {
		rightmost += 1
		if right := ownRow[rightmost]; right != nil {
			surroundingNums = appendUnique(surroundingNums, right)
		}
	}
	// check the row above if exists
	if g.row > 0 {
		aboveRow := (*numbersMatrix)[g.row-1]
		for _, nRef := range aboveRow[leftmost : rightmost+1] {
			if nRef != nil {
				surroundingNums = appendUnique(surroundingNums, nRef)
			}
		}
	}
	// check the row below if exists
	if g.row+1 < len(*numbersMatrix) {
		belowRow := (*numbersMatrix)[g.row+1]
		for _, nRef := range belowRow[leftmost : rightmost+1] {
			if nRef != nil {
				surroundingNums = appendUnique(surroundingNums, nRef)
			}
		}
	}

	if len(surroundingNums) == 2 {
		return *surroundingNums[0] * *surroundingNums[1], true
	}
	return 0, false
}

// append element if it doesn't exists in the slice
func appendUnique(nums []*int, numRef *int) []*int {
	for _, nref := range nums {
		if nref == numRef {
			return nums
		}
	}
	return append(nums, numRef)
}

func Main2() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	rowNum := 0
	allGears := make([]GearPostion, 0, 100)
	numberMatrix := make([][]*int, 0, 100)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums, gears := extractNumbers(line, rowNum)
		numberMatrix = append(numberMatrix, nums)
		allGears = append(allGears, gears...)
		rowNum++
	}
	for _, g := range allGears {
		ratio, exists := g.Ratio(&numberMatrix)
		if exists {
			sum += ratio
		}
	}

	fmt.Printf("Part 2: %d\n", sum)

}

func makeNums() []*int {
	numbers := make([]*int, 0)

	num := new(int)
	for i := 0; i < 10; i++ {
		*num += i
		numbers = append(numbers, num)
	}
	return numbers
}

func printNums(numbers []*int) {
	for _, np := range numbers {
		if np != nil {
			fmt.Printf("%d, ", *np)
		} else {
			fmt.Printf("_, ")
		}
	}
	fmt.Println()
}

func extractNumbers(line string, row int) ([]*int, []GearPostion) {
	numberLocations := make([]*int, len(line))
	currentDigits := make([]rune, 0, 3)
	gears := make([]GearPostion, 0)
	for idx, char := range line {
		if unicode.IsDigit(char) {
			currentDigits = append(currentDigits, char)
		} else {
			if len(currentDigits) > 0 {
				insertNumber(numberLocations, idx-len(currentDigits), currentDigits)
				currentDigits = currentDigits[:0]
			}
			if char == '*' {
				gears = append(gears, GearPostion{row: row, col: idx})
			}
		}
	}
	if len(currentDigits) > 0 {
		insertNumber(numberLocations, len(line)-1-len(currentDigits), currentDigits)
	}
	return numberLocations, gears
}

func makeNumber(digits []rune) *int {
	num, err := strconv.Atoi(string(digits))
	if err != nil {
		log.Fatalf("Couldn't convert '' to number: %v", err)
	}
	return &num
}

func insertNumber(locations []*int, startIndex int, digits []rune) {
	numRef := makeNumber(digits)
	for i := startIndex; i < startIndex+len(digits); i++ {
		locations[i] = numRef
	}
}

package y2024d21

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
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

func parseInput(inputPath string) []string {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(content))
}

type KP_BUTTON int8

const (
	KP_A KP_BUTTON = 0
	KP_0 KP_BUTTON = 1
	KP_1 KP_BUTTON = 2
	KP_2 KP_BUTTON = 3
	KP_3 KP_BUTTON = 4
	KP_4 KP_BUTTON = 5
	KP_5 KP_BUTTON = 6
	KP_6 KP_BUTTON = 7
	KP_7 KP_BUTTON = 8
	KP_8 KP_BUTTON = 9
	KP_9 KP_BUTTON = 10
)

var KP_POSITIONS = [11]vec.Vec2D{
	{X: 3, Y: 2}, // A     0 1 2
	{X: 3, Y: 1}, // 0  0| 7 8 9
	{X: 2, Y: 2}, // 1  1| 6 5 4
	{X: 2, Y: 1}, // 2  2| 3 2 1
	{X: 2, Y: 0}, // 3  3|   0 A
	{X: 1, Y: 2}, // 4
	{X: 1, Y: 1}, // 5
	{X: 1, Y: 0}, // 6
	{X: 0, Y: 0}, // 7
	{X: 0, Y: 1}, // 8
	{X: 0, Y: 2}, // 9
}

func rcButtonPos(rcBtn rune) vec.Vec2D {
	switch rcBtn {
	case '^':
		return vec.Vec2D{X: 0, Y: 1}
	case 'A':
		return vec.Vec2D{X: 0, Y: 2}
	case '<':
		return vec.Vec2D{X: 1, Y: 0}
	case 'v':
		return vec.Vec2D{X: 1, Y: 1}
	case '>':
		return vec.Vec2D{X: 1, Y: 2}
	default:
		msg := fmt.Sprintf("Unknown RC button: %c\n", rcBtn)
		panic(msg)
	}
}

func kpButtonPos(kpBtn rune) vec.Vec2D {
	switch kpBtn {
	case 'A':
		return vec.Vec2D{X: 3, Y: 2}
	case '0':
		return vec.Vec2D{X: 3, Y: 1}
	case '1':
		return vec.Vec2D{X: 2, Y: 0}
	case '2':
		return vec.Vec2D{X: 2, Y: 1}
	case '3':
		return vec.Vec2D{X: 2, Y: 2}
	case '4':
		return vec.Vec2D{X: 1, Y: 0}
	case '5':
		return vec.Vec2D{X: 1, Y: 1}
	case '6':
		return vec.Vec2D{X: 1, Y: 2}
	case '7':
		return vec.Vec2D{X: 0, Y: 0}
	case '8':
		return vec.Vec2D{X: 0, Y: 1}
	case '9':
		return vec.Vec2D{X: 0, Y: 2}
	default:
		msg := fmt.Sprintf("unknown button: %c\n", kpBtn)
		panic(msg)
	}
}

func pathToRcButton(from, to rune) []rune {
	return pathToButton(rcButtonPos(from), rcButtonPos(to), true)
}

func pathToKpButton(from, to rune) []rune {
	// fmt.Printf("[%q] -> [%q]\n", from, to)
	return pathToButton(kpButtonPos(from), kpButtonPos(to), false)
}

func pathToButton(fromPos, toPos vec.Vec2D, isRcPad bool) []rune {
	xDist := toPos.X - fromPos.X
	yDist := toPos.Y - fromPos.Y

	var vertMove, horizMove rune
	if xDist > 0 {
		vertMove = 'v'
	} else {
		vertMove = '^'
	}
	if yDist > 0 {
		horizMove = '>'
	} else {
		horizMove = '<'
	}

	moves := make([]rune, Abs(xDist)+Abs(yDist)+1)
	firstMoveVert := true
	if !isRcPad && vertMove == 'v' && horizMove == '>' {
		firstMoveVert = false
	} else if isRcPad && vertMove == '^' && horizMove == '>' {
		firstMoveVert = false
	}

	for i := range moves {
		if (firstMoveVert && i < Abs(xDist)) || (!firstMoveVert && i >= Abs(yDist)) {
			moves[i] = vertMove
		} else {
			moves[i] = horizMove
		}
	}
	moves[len(moves)-1] = 'A'
	return moves
}

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// func translateRcToRc(buttons []rune, layers int) []rune {
// 	// fmt.Printf("translateRcToRc(%q, %d)\n", string(buttons), layers)
// 	for i := 0; i < layers; i++ {
// 		fromBtn := 'A'
// 		var newButtons []rune
// 		for _, btn := range buttons {
// 			path := pathToRcButton(fromBtn, btn)
// 			newButtons = append(newButtons, path...)
// 			fromBtn = btn
// 		}
// 		// fmt.Printf("%q -> %q\n", string(buttons), string(newButtons))
// 		buttons = newButtons
// 	}
// 	return buttons
// }

func translateRcToRc(buttons []rune, layers int) []rune {
	fmt.Printf("%s (%d)\n", string(buttons), layers)
	if layers == 0 {
		return buttons
	}
	fromBtn := 'A'
	var newButtons []rune
	for _, btn := range buttons {
		path := pathToRcButton(fromBtn, btn)
		newButtons = append(newButtons, path...)
		fromBtn = btn
	}
	return translateRcToRc(newButtons, layers-1)
}

func Part1(codes []string) int {
	// KeyPad <- RC <- RC <- RC

	// 379A
	// ^A|^^<<A|>>A|vvvA

	// ^A|^^<<A|>>A|vvvA|

	total := 0
	for _, code := range codes {
		if code != "379A" {
			continue
		}
		fmt.Println()
		fromBtn := 'A'
		pathLen := 0
		for _, btn := range code {
			path := pathToKpButton(fromBtn, btn)

			fromBtn = btn
			translatedPath := translateRcToRc(path, 2)
			fmt.Println()
			// fmt.Printf("%s|", string(path))
			pathLen += len(translatedPath)
		}
		// fmt.Println()
		complexity := computeCodeComplexity(code, pathLen)
		fmt.Printf("%q -> %d\n", code, complexity)
		total += complexity
	}
	return total
}

func Part2(codes []string) int {
	return 0
}

func computeCodeComplexity(code string, length int) int {
	digits := make([]rune, 0, len(code))
	for _, ch := range code {
		if unicode.IsDigit(ch) {
			digits = append(digits, ch)
		}
	}

	n, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}
	fmt.Printf("complexity of %q: %d x %d == %d\n", code, length, n, n*length)
	return n * length
}

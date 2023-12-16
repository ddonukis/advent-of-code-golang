package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (item T, ok bool) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		ok = true
	}

	return item, ok
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func (s *Stack[T]) String() string {
	parts := make([]string, len(s.items)+1)
	for i, item := range s.items {
		parts[i] = fmt.Sprintf("- %v", item)
	}
	parts[len(parts)-1] = "Stack:"
	slices.Reverse(parts)
	return strings.Join(parts, "\n")
}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	contraption := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		contraption = append(contraption, []rune(line))
	}

	tracer := InitTraceWithCache()

	eP1, _ := calcEnergizedTiles(Beam{0, 0, DirEast}, contraption, tracer)
	fmt.Printf("Part 1 energized: %d\n", eP1)

	contraptionW := len(contraption[0])
	contraptionH := len(contraption)

	startingBeams := make([]Beam, 0, 4+contraptionH*contraptionW)
	// corners
	startingBeams = append(startingBeams, Beam{0, 0, DirEast})
	startingBeams = append(startingBeams, Beam{0, 0, DirSouth})
	startingBeams = append(startingBeams, Beam{contraptionW - 1, 0, DirWest})
	startingBeams = append(startingBeams, Beam{contraptionW - 1, 0, DirSouth})
	startingBeams = append(startingBeams, Beam{contraptionW - 1, contraptionH - 1, DirNorth})
	startingBeams = append(startingBeams, Beam{contraptionW - 1, contraptionH - 1, DirWest})
	startingBeams = append(startingBeams, Beam{0, contraptionH - 1, DirEast})
	startingBeams = append(startingBeams, Beam{0, contraptionH - 1, DirNorth})
	// north edge
	for i := 1; i < contraptionW-1; i++ {
		startingBeams = append(startingBeams, Beam{X: i, Y: 0, direction: DirSouth})
	}
	// south edge
	for i := 1; i < contraptionW-1; i++ {
		startingBeams = append(startingBeams, Beam{X: i, Y: contraptionH - 1, direction: DirNorth})
	}
	// west edge
	for i := 1; i < contraptionH-1; i++ {
		startingBeams = append(startingBeams, Beam{X: 0, Y: i, direction: DirEast})
	}
	// east edge
	for i := 1; i < contraptionH-1; i++ {
		startingBeams = append(startingBeams, Beam{X: contraptionW - 1, Y: i, direction: DirWest})
	}

	maxEnergized := 0
	var bestStartingBeam Beam
	for _, beam := range startingBeams {
		e, _ := calcEnergizedTiles(beam, contraption, tracer)
		if e > maxEnergized {
			maxEnergized = e
			bestStartingBeam = beam
		}
	}
	fmt.Printf("Part 2: Best beam %v -> %d energized\n", bestStartingBeam, maxEnergized)

}

func calcEnergizedTiles(
	startingBeam Beam,
	contraption [][]rune,
	tracer func(Beam, [][]rune) ([][]bool, []Beam),
) (energized, iterations int) {
	beams := NewStack[Beam]()
	beams.Push(startingBeam)

	var totalEnergized [][]bool

	totalIterations := 0
	traced := make(map[Beam]bool)
	for !beams.IsEmpty() {
		curBeam, ok := beams.Pop()
		if !ok {
			break
		}
		energizedMap, newBeams := tracer(curBeam, contraption)
		traced[curBeam] = true
		updateTotalEnergizedMap(&totalEnergized, energizedMap)

		for _, b := range newBeams {
			if !traced[b] {
				beams.Push(b)
			}
		}
		totalIterations++
		if totalIterations > 1000 {
			break
		}
	}
	energizedCount := 0
	for _, ln := range totalEnergized {
		for _, tile := range ln {
			if tile {
				energizedCount++
			} else {
			}
		}
	}
	return energizedCount, totalIterations
}

type Direction struct {
	X, Y int
}

var (
	DirNorth Direction = Direction{0, -1}
	DirSouth           = Direction{0, 1}
	DirWest            = Direction{-1, 0}
	DirEast            = Direction{1, 0}
)

type Beam struct {
	X, Y      int
	direction Direction
}

func (beam *Beam) Move() {
	beam.X += beam.direction.X
	beam.Y += beam.direction.Y
}

// trace one beam until it goes out of bounds or gets split. Return energized tiles map and a slice
// of beams it got split into.
func trace(beam Beam, contraption [][]rune) (energizedTiles [][]bool, newBeams []Beam) {
	energizedMap := make([][]bool, len(contraption))
	for rowIdx := range energizedMap {
		energizedMap[rowIdx] = make([]bool, len(contraption[rowIdx]))
	}
	mapWidth, mapHeight := len(contraption[0]), len(contraption)

	newBeams = make([]Beam, 0, 2)

beamTracing:
	for inBounds(beam, mapWidth, mapHeight) {
		energizedMap[beam.Y][beam.X] = true
		switch contraption[beam.Y][beam.X] {
		case '/':
			switch beam.direction {
			case DirNorth:
				beam.direction = DirEast
			case DirSouth:
				beam.direction = DirWest
			case DirWest:
				beam.direction = DirSouth
			case DirEast:
				beam.direction = DirNorth
			default:
				log.Fatalf("Unexpected direction: %v\n", beam.direction)
			}
		case '\\':
			switch beam.direction {
			case DirNorth:
				beam.direction = DirWest
			case DirSouth:
				beam.direction = DirEast
			case DirWest:
				beam.direction = DirNorth
			case DirEast:
				beam.direction = DirSouth
			default:
				log.Fatalf("Unexpected direction: %v\n", beam.direction)
			}
		case '-':
			if beam.direction == DirNorth || beam.direction == DirSouth {
				eastBeam := Beam{X: beam.X, Y: beam.Y, direction: DirEast}
				eastBeam.Move()
				if inBounds(eastBeam, mapWidth, mapHeight) {
					newBeams = append(newBeams, eastBeam)
				}
				westBeam := Beam{X: beam.X, Y: beam.Y, direction: DirWest}
				westBeam.Move()
				if inBounds(westBeam, mapWidth, mapHeight) {
					newBeams = append(newBeams, westBeam)
				}
				break beamTracing
			}
		case '|':
			if beam.direction == DirEast || beam.direction == DirWest {
				northBeam := Beam{X: beam.X, Y: beam.Y, direction: DirNorth}
				northBeam.Move()
				if inBounds(northBeam, mapWidth, mapHeight) {
					newBeams = append(newBeams, northBeam)
				}
				southBeam := Beam{X: beam.X, Y: beam.Y, direction: DirSouth}
				southBeam.Move()
				if inBounds(southBeam, mapWidth, mapHeight) {
					newBeams = append(newBeams, southBeam)
				}
				break beamTracing
			}
		default:
		}
		beam.Move()
	}
	return energizedMap, newBeams
}

type traceOutput struct {
	energizedTiles *[][]bool
	newBeams       *[]Beam
}

func InitTraceWithCache() func(Beam, [][]rune) ([][]bool, []Beam) {
	cache := make(map[Beam]traceOutput)
	return func(beam Beam, contraption [][]rune) (energizedTiles [][]bool, newBeams []Beam) {
		output, cacheExists := cache[beam]
		if cacheExists {
			return *output.energizedTiles, *output.newBeams
		}
		energizedTiles, newBeams = trace(beam, contraption)
		cache[beam] = traceOutput{&energizedTiles, &newBeams}
		return energizedTiles, newBeams
	}
}

func inBounds(beam Beam, width, height int) bool {
	return (beam.X >= 0 && beam.X < width) && (beam.Y >= 0 && beam.Y < height)
}

func updateTotalEnergizedMap(totalEnergized *[][]bool, energizedMap [][]bool) {
	if *totalEnergized == nil {
		*totalEnergized = energizedMap
	} else {
		for rowIdx := range *totalEnergized {
			for colIdx := range (*totalEnergized)[rowIdx] {
				(*totalEnergized)[rowIdx][colIdx] = (*totalEnergized)[rowIdx][colIdx] || energizedMap[rowIdx][colIdx]
			}
		}
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	id     int
	keys   map[string]int
	values []int
}

func (b *Box) Add(label string, focalLength int) {
	index, exists := b.keys[label]
	if exists {
		b.values[index] = focalLength
	} else {
		index = len(b.values)
		b.values = append(b.values, focalLength)
		b.keys[label] = index
	}
}

func (b *Box) Remove(label string) {
	index, exists := b.keys[label]
	if exists {
		b.values[index] = -1
		delete(b.keys, label)
	}
}

func (b *Box) String() string {
	strParts := make([]string, 1, len(b.keys)+1)
	strParts[0] = fmt.Sprintf("Box %d:", b.id)
	for k, idx := range b.keys {
		p := fmt.Sprintf("[%s %d]", k, b.values[idx])
		strParts = append(strParts, p)
	}
	return strings.Join(strParts, " ")
}

func (b *Box) FocusingPower() int {
	fp := 0
	boxNo := b.id + 1
	slotNo := 0
	for _, focalLength := range b.values {
		if focalLength > 0 {
			slotNo++
			fp += boxNo * slotNo * focalLength
		}
	}
	return fp
}

func NewBox(id int) *Box {
	return &Box{id: id, keys: make(map[string]int), values: make([]int, 0, 1)}
}

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sequenceScanner := bufio.NewScanner(f)
	sequenceScanner.Split(splitSequence)

	total := 0
	storage := make([]*Box, 256)
	totalFocusingPower := 0
	for sequenceScanner.Scan() {
		instruction := sequenceScanner.Text()
		instructionHash := hash(instruction)
		total += instructionHash

		execute(instruction, storage)
	}
	for _, b := range storage {
		if b != nil {
			fmt.Println(b)
			totalFocusingPower += b.FocusingPower()
		}
	}
	fmt.Printf("Hash sum (P1): %d\n", total)
	fmt.Printf("Focusing power (P2): %d\n", totalFocusingPower)

}

func execute(instruction string, storage []*Box) {
	if idx := strings.Index(instruction, "-"); idx > 0 {
		label := instruction[:idx]
		boxId := hash(label)
		if b := storage[boxId]; b != nil {
			b.Remove(label)
		}
	} else {
		labelAndFocalLength := strings.Split(instruction, "=")
		boxId := hash(labelAndFocalLength[0])
		focalLength, err := strconv.Atoi(labelAndFocalLength[1])
		if err != nil {
			log.Fatalf("Cannot parse instruction '%s': %v\n", instruction, err)
		}
		b := storage[boxId]
		if b == nil {
			b = NewBox(boxId)
			storage[boxId] = b
		}
		b.Add(labelAndFocalLength[0], focalLength)
	}
}

func hash(s string) (value int) {
	for _, ch := range s {
		value += int(ch)
		value *= 17
		value = value % 256
	}
	return
}

func splitSequence(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, trimEnd(data[0:i]), nil
	}
	if atEOF {
		return len(data), trimEnd(data), nil
	}
	return 0, nil, nil
}

func trimEnd(data []byte) (trimmedData []byte) {
	trimmedData = data
	switch {
	case len(data) > 1 && slices.Equal(data[len(data)-2:], []byte{'\r', '\n'}):
		trimmedData = data[:len(data)-2]
	case len(data) > 0 && data[len(data)-1] == '\n':
		trimmedData = data[:len(data)-1]

	}
	return trimmedData
}

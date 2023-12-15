package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sequenceScanner := bufio.NewScanner(f)
	sequenceScanner.Split(splitSequence)
	total := 0
	for sequenceScanner.Scan() {
		token := sequenceScanner.Text()
		tokenHash := hash(token)
		fmt.Printf("'%s' -> %d\n", sequenceScanner.Text(), tokenHash)
		total += tokenHash
	}
	fmt.Printf("Hash sum (P1): %d\n", total)

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

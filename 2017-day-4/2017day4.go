package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	counter := 0
	counterP2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		uniqueWords := ToSet(words)

		if len(words) == uniqueWords.Len() {
			counter++
		} else {
			continue
		}
		uniqueSortedChars := NewSet[string]()
		for _, word := range words {
			uniqueSortedChars.Add(sortChars(word))
		}
		if len(words) == uniqueSortedChars.Len() {
			counterP2++
		}
	}
	fmt.Printf("Part 1: %d\n", counter)
	fmt.Printf("Part 2: %d\n", counterP2)
}

func sortChars(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func ToSet[T string](s []T) *Set[T] {
	items := make(map[T]bool)
	for _, item := range s {
		items[item] = true
	}
	return &Set[T]{items: items}
}

func NewSet[T string]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

type Set[T string] struct {
	items map[T]bool
}

func (set *Set[T]) Add(item T) {
	set.items[item] = true
}

func (set *Set[T]) Len() int {
	return len(set.items)
}

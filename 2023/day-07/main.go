package main

import "fmt"

var cardRanking = [...]rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

type Hand [5]rune

type HandWithBid struct {
	hand Hand
	bid  int
}

func (hwb HandWithBid) String() string {
	return fmt.Sprintf("{%s - %d}", string(hwb.hand[:]), hwb.bid)
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	const path = "input.txt"

	part1Answer := part1(path)
	part2Answer := part2(path)

	fmt.Printf("\nPart 1 winnings: %d\n", part1Answer)
	fmt.Printf("\nPart 2 winnings: %d\n", part2Answer)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
)

var cardRankingP2 = [...]rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func handTypeP2(hand Hand) HandType {
	uniqueCardCount := make(map[rune]int)
	for _, card := range hand {
		val, exists := uniqueCardCount[card]
		if exists {
			uniqueCardCount[card] = val + 1
		} else {
			uniqueCardCount[card] = 1
		}
	}
	jokerCount, jokerInHand := uniqueCardCount['J']
	if jokerInHand {
		return differentiateHandTypesWithJoker(uniqueCardCount, jokerCount)
	}

	return differentiateHandTypes(uniqueCardCount)
}

func differentiateHandTypesWithJoker(uniqueCardCount map[rune]int, jokerCount int) HandType {
	switch len(uniqueCardCount) {
	case 5:
		return OnePair
	case 4:
		return ThreeOfAKind
	case 3:
		// JJJ23
		// JJ223 JJ233
		// J2333 J2233 J2223
		if jokerCount == 2 || jokerCount == 3 {
			return FourOfAKind
		}
		for _, count := range uniqueCardCount {
			if count == 3 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 2, 1:
		return FiveOfAKind
	}
	return OnePair
}

func differentiateHandTypes(uniqueCardCount map[rune]int) HandType {
	switch len(uniqueCardCount) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, count := range uniqueCardCount {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 2:
		for _, count := range uniqueCardCount {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 1:
		return FiveOfAKind
	}
	return HighCard
}

func handValueP2(hand Hand) string {
	cardValues := make([]int, 5)
	for i, card := range hand {
		cardVal := slices.Index(cardRankingP2[:], card)
		cardValues[i] = cardVal
	}
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", handTypeP2(hand), cardValues[0], cardValues[1], cardValues[2], cardValues[3], cardValues[4])
}

func part2(filepath string) int {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	handWithBids := make([]HandWithBid, 0)
	for scanner.Scan() {
		line := scanner.Text()
		hwb, err := parseRow(line)
		if err == nil {
			handWithBids = append(handWithBids, hwb)
		} else {
			log.Fatalf("Couldn't parse line '%s': %v\n", line, err)
		}
		fmt.Printf("%s: %d: %s\n", hwb, handType(hwb.hand), handValue(hwb.hand))
	}
	sort.Slice(handWithBids, func(i, j int) bool {
		return handValueP2(handWithBids[i].hand) < handValueP2(handWithBids[j].hand)
	})

	winnings := 0
	for idx, hwb := range handWithBids {
		winnings += (idx + 1) * hwb.bid
	}

	return winnings
}

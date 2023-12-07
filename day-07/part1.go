package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func handType(hand Hand) HandType {
	uniqueCardCount := make(map[rune]int)
	for _, card := range hand {
		val, exists := uniqueCardCount[card]
		if exists {
			uniqueCardCount[card] = val + 1
		} else {
			uniqueCardCount[card] = 1
		}
	}
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

func handValue(hand Hand) string {
	cardValues := make([]int, 5)
	for i, card := range hand {
		cardVal := slices.Index(cardRanking[:], card)
		cardValues[i] = cardVal
	}
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", handType(hand), cardValues[0], cardValues[1], cardValues[2], cardValues[3], cardValues[4])
}

func parseRow(line string) (HandWithBid, error) {
	split := strings.Fields(strings.TrimSpace(line))
	hwb := HandWithBid{}
	if len(split[0]) != 5 {
		return hwb, fmt.Errorf("invalid hand '%s'", split[0])
	}
	for i, ch := range split[0] {
		hwb.hand[i] = ch
	}
	bid, err := strconv.Atoi(split[1])
	if err == nil {
		hwb.bid = bid
	} else {
		log.Printf("Cannot convert '%s' to int\n", split[1])
		return hwb, err
	}
	return hwb, nil
}

func part1(filepath string) int {
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
		return handValue(handWithBids[i].hand) < handValue(handWithBids[j].hand)
	})

	winnings := 0
	for idx, hwb := range handWithBids {
		winnings += (idx + 1) * hwb.bid
	}

	return winnings
}

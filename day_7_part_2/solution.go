package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard     = iota
	onePair      = iota
	twoPair      = iota
	threeOfAKind = iota
	fullHouse    = iota
	fourOfAKind  = iota
	fiveOfAKind  = iota
)

type Hand struct {
	cards    string
	bid      int
	handType int
}

type By func(p1, p2 *Hand) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(hands []Hand) {
	hs := &handSorter{
		hands: hands,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(hs)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type handSorter struct {
	hands []Hand
	by    func(h1, h2 *Hand) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *handSorter) Len() int {
	return len(s.hands)
}

// Swap is part of sort.Interface.
func (s *handSorter) Swap(i, j int) {
	s.hands[i], s.hands[j] = s.hands[j], s.hands[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *handSorter) Less(i, j int) bool {
	return s.by(&s.hands[i], &s.hands[j])
}

func main() {
	exampleFlagPtr := flag.Bool("example", false, "Use the example file if set")
	flag.Parse()
	// Read input file
	inputFile := "input.txt"
	if *exampleFlagPtr {
		inputFile = "example.txt"
	}
	content, _ := os.ReadFile(inputFile)
	content_string := string(content)
	output := 0
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// Parse hands and determine type
	hands := make([]Hand, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		handType := determineHandType(parts[0])
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{
			cards:    parts[0],
			bid:      bid,
			handType: handType,
		})
	}
	// Sort
	value := func(h1, h2 *Hand) bool {
		if h1.handType != h2.handType {
			return h1.handType < h2.handType
		}
		for i := range h1.cards {
			h1Number := convertToValue(h1.cards[i])
			h2Number := convertToValue(h2.cards[i])
			if h1Number != h2Number {
				return h1Number < h2Number
			}
		}
		return true
	}
	By(value).Sort(hands)
	// Calculate output
	for rank, hand := range hands {
		output += hand.bid * (rank + 1)
	}
	// Print the result
	fmt.Println(output)
}

func convertToValue(character byte) int {
	if value, err := strconv.Atoi(string(character)); err == nil {
		return value
	}
	valueMap := map[byte]int{
		'T': 10,
		'J': 1,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return valueMap[character]
}

func determineHandType(hand string) int {
	counts := make(map[rune]int, 0)
	for _, letter := range hand {
		if _, exists := counts[letter]; !exists {
			counts[letter] = 1
		} else {
			counts[letter]++
		}
	}
	jokerCount := counts['J']
	if len(counts) == 1 {
		return fiveOfAKind
	} else if len(counts) == 2 {
		minCount := 5
		for _, count := range counts {
			if count < minCount {
				minCount = count
			}
		}
		if minCount > 1 {
			if jokerCount >= 1 {
				return fiveOfAKind
			}
			return fullHouse
		} else {
			if jokerCount >= 1 {
				return fiveOfAKind
			}
			return fourOfAKind
		}
	} else if len(counts) == 3 {
		maxCount := 0
		for _, count := range counts {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount > 2 {
			if jokerCount == 1 || jokerCount == 3 {
				return fourOfAKind
			}
			return threeOfAKind
		} else {
			if jokerCount == 1 {
				return fullHouse
			} else if jokerCount == 2 {
				return fourOfAKind
			}
			return twoPair
		}
	} else if len(counts) == 4 {
		if jokerCount > 0 {
			return threeOfAKind
		}
		return onePair
	}
	if jokerCount > 0 {
		return onePair
	}
	return highCard
}

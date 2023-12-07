package main

import (
	"aoc2023/aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var sample = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var Labels = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var Labels2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type Hand struct {
	Cards []string
	Bid   int
	Type  int
}

func getHandType(cards []string) int {
	have := make(map[string]int)
	for _, c := range cards {
		have[c]++
	}

	if len(have) == 1 {
		return FiveOfKind
	} else if len(have) == 5 {
		return HighCard
	}

	var (
		four   bool
		three  bool
		two    bool
		twotwo bool
	)
	for _, n := range have {
		if n == 4 {
			four = true
		}
		if n == 3 {
			three = true
		}
		if two == true && n == 2 {
			twotwo = true
		}
		if n == 2 {
			two = true
		}
	}

	if four {
		return FourOfKind
	} else if three && two {
		return FullHouse
	} else if three {
		return ThreeOfKind
	} else if twotwo {
		return TwoPair
	} else if two {
		return OnePair
	}

	return 0
}

func getHands(input string) []Hand {
	var hands []Hand
	for _, l := range strings.Split(input, "\n") {
		var temphand Hand
		split := strings.Split(l, " ")

		temphand.Cards = strings.Split(split[0], "")
		temphand.Bid, _ = strconv.Atoi(split[1])
		temphand.Type = getHandType(temphand.Cards)

		hands = append(hands, temphand)
	}
	return hands
}

func first(input string) {
	hands := getHands(input)

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for ii, c := range hands[i].Cards {
				if c != hands[j].Cards[ii] {
					return Labels[c] < Labels[hands[j].Cards[ii]]
				}
			}
		}

		return hands[i].Type < hands[j].Type
	})

	sum := 0
	for i, h := range hands {
		sum += h.Bid * (i + 1)
	}
	fmt.Println("total winnings:", sum)
}

func getHandType2(cards []string) int {
	have := make(map[string]int)
	for _, c := range cards {
		have[c]++
	}

	if have["J"] > 0 {
		most := ""
		mostc := 0
		for c, v := range have {
			if c != "J" && v > mostc {
				most = c
				mostc = v
			}
		}
		have[most] += have["J"]
		delete(have, "J")
	}

	if len(have) == 1 {
		return FiveOfKind
	} else if len(have) == 5 {
		return HighCard
	}

	var (
		four   bool
		three  bool
		two    bool
		twotwo bool
	)
	for _, n := range have {
		if n == 4 {
			four = true
		}
		if n == 3 {
			three = true
		}
		if two == true && n == 2 {
			twotwo = true
		}
		if n == 2 {
			two = true
		}
	}

	if four {
		return FourOfKind
	} else if three && two {
		return FullHouse
	} else if three {
		return ThreeOfKind
	} else if twotwo {
		return TwoPair
	} else if two {
		return OnePair
	}

	return 0
}

func getHands2(input string) []Hand {
	var hands []Hand
	for _, l := range strings.Split(input, "\n") {
		var temphand Hand
		split := strings.Split(l, " ")

		temphand.Cards = strings.Split(split[0], "")
		temphand.Bid, _ = strconv.Atoi(split[1])
		temphand.Type = getHandType2(temphand.Cards)

		hands = append(hands, temphand)
	}
	return hands
}

func second(input string) {
	hands := getHands2(input)

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for ii, c := range hands[i].Cards {
				if c != hands[j].Cards[ii] {
					return Labels2[c] < Labels2[hands[j].Cards[ii]]
				}
			}
		}

		return hands[i].Type < hands[j].Type
	})

	sum := 0
	for i, h := range hands {
		sum += h.Bid * (i + 1)
	}
	fmt.Println("total winnings:", sum)
}

func main() {
	//input := sample
	input := aoc.InputForDay("07")

	first(input)
	second(input)
}

package main

import (
	"fmt"
	"strings"
)

var sample = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

type Card struct {
	ID      string
	Nums    []string
	Winners []string
}

func getCards(input string) []Card {
	var cards []Card

	for _, c := range strings.Split(input, "\n") {
		var tempcard Card

		split := strings.Split(c, ": ")
		tempcard.ID = strings.TrimPrefix(split[0], "Card ")

		split = strings.Split(split[1], " | ")
		for _, n := range strings.Split(split[0], " ") {
			n = strings.TrimSpace(n)
			if n != "" {
				tempcard.Winners = append(tempcard.Winners, n)
			}
		}

		for _, n := range strings.Split(split[1], " ") {
			n = strings.TrimSpace(n)
			if n != "" {
				tempcard.Nums = append(tempcard.Nums, n)
			}
		}
		cards = append(cards, tempcard)
	}

	return cards
}

func isIn(thing string, arr []string) bool {
	for _, a := range arr {
		if a == thing {
			return true
		}
	}
	return false
}

func first(input string) {
	cards := getCards(input)

	sum := 0
	for _, c := range cards {
		val := 0
		for _, num := range c.Nums {
			if isIn(num, c.Winners) {
				if val == 0 {
					val = 1
				} else {
					val = val * 2
				}
			}
		}
		sum += val
		fmt.Println("card:", c.ID, "value:", val, "new sum:", sum)
	}
}

func cardLoop(vals []int) int {
	var newvals []int
	for i := 0; i < 20; i++ {
		for n, v := range vals {
			for i := 0; i < v; i++ {
				newvals = append(newvals, vals[n+i+1])
			}
			break
		}
		fmt.Println(vals, newvals)
		vals = newvals
	}

	return len(newvals)
}

func second(input string) {
	cards := getCards(input)

	var vals []int
	for _, c := range cards {
		val := 0
		for _, num := range c.Nums {
			if isIn(num, c.Winners) {
				val++
			}
		}
		vals = append(vals, val)
	}

	d := cardLoop(vals)

	fmt.Println(d)
}

func main() {
	input := sample
	//input := aoc.InputForDay("04")

	//first(input)
	second(input)
}
